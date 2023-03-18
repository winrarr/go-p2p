package peer

import (
	"log"
	"net"
	"p2p/rpc"
)

type Peer struct {
	conn       *net.UDPConn
	knownPeers []*net.UDPAddr
}

func CreateGenesisPeer(lPort int) Peer {
	conn := newConnection(lPort)
	p := Peer{
		conn: conn,
	}
	go p.readFromConnection(conn)

	return p
}

func CreatePeerAndConnect(lPort int, rIp string, rPort int) Peer {
	conn := newConnection(lPort)
	rAddr := &net.UDPAddr{
		IP:   net.ParseIP(rIp),
		Port: rPort,
	}
	p := Peer{
		conn:       conn,
		knownPeers: []*net.UDPAddr{rAddr},
	}
	go p.readFromConnection(conn)
	p.AskForConnections(rAddr)
	return p
}

func newConnection(lPort int) *net.UDPConn {
	lAddr := &net.UDPAddr{
		Port: lPort,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal()
	}
	return conn
}

func (p *Peer) readFromConnection(conn *net.UDPConn) {
	rpc := rpc.NewRpc(rpc.DefaultUDPReadWriter(256))
	rpc.RegisterProcedure("sendConnections", p.sendConnections)
	rpc.RegisterProcedure("connections", p.connections)
	rpc.Start(conn)
}

func (p *Peer) AskForConnections(addr *net.UDPAddr) {
	p.sendTo("sendConnections ", addr)
	buf := make([]byte, 2048)
	_, _, err := p.conn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal()
	}
	println(string(buf))
}

func (p *Peer) SendTooAll(message string) {
	for _, addr := range p.knownPeers {
		p.sendTo(message, addr)
	}
}

func (p *Peer) sendTo(message string, addr *net.UDPAddr) {
	_, err := p.conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Peer) Address() net.Addr {
	return p.conn.LocalAddr()
}
