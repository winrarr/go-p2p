package rpc

import (
	"errors"
	"log"
	"net"
	"strings"
)

type Responder[T net.Conn] func([]byte) ([]byte, error)
type readWriter[T net.Conn] func(T, Responder[T])
type procedure[T net.Conn] func(...string) []string

type rpc[T net.Conn] struct {
	rw         readWriter[T]
	procedures map[string]procedure[T]
}

func NewRpc[T net.Conn](rw readWriter[T]) rpc[T] {
	return rpc[T]{
		rw:         rw,
		procedures: map[string]procedure[T]{},
	}
}

func (r *rpc[T]) RegisterProcedure(command string, p procedure[T]) {
	r.procedures[command] = p
}

func (r *rpc[T]) UnregisterProcedure(command string) {
	delete(r.procedures, command)
}

func (r *rpc[T]) responder(request []byte) ([]byte, error) {
	fp := strings.Split(string(request), " ")
	if _, ok := r.procedures[fp[0]]; !ok {
		return nil, errors.New("no such procedure found")
	}
	procedure, payload := r.procedures[fp[0]], fp[1:]
	response := procedure(payload...)
	return []byte(strings.Join(response, " ")), nil
}

func (r *rpc[T]) Start(conn T) {
	for {
		r.rw(conn, r.responder)
	}
}

func DefaultUDPReadWriter(maxBytes int) readWriter[*net.UDPConn] {
	return func(conn *net.UDPConn, respond Responder[*net.UDPConn]) {
		buf := make([]byte, 256)
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal()
		}
		response, err := respond(buf)
		if err == nil {
			conn.WriteToUDP(response, addr)
		}
	}
}
