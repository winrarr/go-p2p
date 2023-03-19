package rpc

import (
	"encoding/json"
	"log"
	"net"
)

type procedure func(*Request)
type reader func() *Request
type writer func([]byte, *net.UDPAddr)

type Rpc struct {
	conn       *net.UDPConn
	procedures map[string]procedure
	read       reader
	write      writer
}

type Request struct {
	rpc       *Rpc
	Addr      *net.UDPAddr
	procedure string
	Payload   []byte
}

type requestData struct {
	Procedure    string
	PayloadBytes []byte
}

func NewRpc(conn *net.UDPConn) *Rpc {
	rpc := Rpc{
		conn:       conn,
		procedures: map[string]procedure{},
	}

	buf := make([]byte, 1024)
	reader := func() *Request {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("failed to read from udp")
		}
		// println("received:", string(buf))
		var data requestData
		err = json.Unmarshal(buf[:n], &data)
		if err != nil {
			log.Fatal("error unmarshalling data")
		}

		return &Request{
			rpc:       &rpc,
			Addr:      addr,
			procedure: data.Procedure,
			Payload:   data.PayloadBytes,
		}
	}

	writer := func(bytes []byte, addr *net.UDPAddr) {
		_, err := conn.WriteToUDP(bytes, addr)
		if err != nil {
			log.Fatal("failed to write to udp")
		}
		// println("sent:", string(bytes))
	}

	rpc.read = reader
	rpc.write = writer

	return &rpc
}

func (r *Request) Respond(procedure string, payload any) {
	r.rpc.SendTo(procedure, payload, r.Addr)
}

func (r *Rpc) RegisterProcedure(command string, p procedure) {
	r.procedures[command] = p
}

func (r *Rpc) UnregisterProcedure(command string) {
	delete(r.procedures, command)
}

func (r *Rpc) Listen() {
	for {
		req := r.read()
		// println(r.conn.LocalAddr().String(), "received procedure", req.procedure)
		procedure, ok := r.procedures[req.procedure]
		if ok {
			procedure(req)
		}
	}
}

func (r *Rpc) SendTo(procedure string, payload any, addr *net.UDPAddr) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("error marshalling the payload")
	}
	data := requestData{
		Procedure:    procedure,
		PayloadBytes: payloadBytes,
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal("error marshalling the request")
	}
	r.write(bytes, addr)
	// println(r.conn.LocalAddr().String(), "sent procedure", data.Procedure, "to", addr.String())
}
