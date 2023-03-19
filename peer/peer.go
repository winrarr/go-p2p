package peer

import (
	"log"
	"net"
	"p2p/rpc"
)

type Peer struct {
	rpc        *rpc.Rpc
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
	p.SendHello(rAddr)
	p.SendSendConnections(rAddr)
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
	rpc := rpc.NewRpc(conn)
	rpc.RegisterProcedure("sendConnections", p.getSendConnections)
	rpc.RegisterProcedure("connections", p.getConnections)
	rpc.RegisterProcedure("hello", p.getHello)
	rpc.Listen()
}

func (p *Peer) SendTooAll(procedure string, payload any) {
	for _, addr := range p.knownPeers {
		p.rpc.SendTo(procedure, payload, addr)
	}
}

func (p *Peer) Address() net.Addr {
	return p.conn.LocalAddr()
}

func (p *Peer) KnownPeers() []*net.UDPAddr {
	return p.knownPeers
}
