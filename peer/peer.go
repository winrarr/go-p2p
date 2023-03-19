package peer

import (
	"log"
	"net"
	"p2p/rpc"
)

type Peer struct {
	Name       string
	rpc        *rpc.Rpc
	conn       *net.UDPConn
	knownPeers []*net.UDPAddr
}

func CreateGenesisPeer(name string, lPort int) Peer {
	conn, lAddr := newConnection(lPort)
	p := Peer{
		Name:       name,
		conn:       conn,
		knownPeers: []*net.UDPAddr{lAddr},
	}
	conn.LocalAddr()
	go p.readFromConnection(conn)

	return p
}

func CreatePeerAndConnect(name string, lPort int, rIp string, rPort int) Peer {
	conn, lAddr := newConnection(lPort)
	rAddr := &net.UDPAddr{
		IP:   net.ParseIP(rIp),
		Port: rPort,
	}
	p := Peer{
		Name:       name,
		conn:       conn,
		knownPeers: []*net.UDPAddr{lAddr, rAddr},
	}
	p.readFromConnection(conn)
	p.SendHello(rAddr)
	p.SendSendConnections(rAddr)
	return p
}

func newConnection(lPort int) (*net.UDPConn, *net.UDPAddr) {
	lAddr := &net.UDPAddr{
		Port: lPort,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal()
	}
	return conn, lAddr
}

func (p *Peer) readFromConnection(conn *net.UDPConn) {
	p.rpc = rpc.NewRpc(conn, 1024)
	p.rpc.RegisterProcedure("sendConnections", p.getSendConnections)
	p.rpc.RegisterProcedure("connections", p.getConnections)
	p.rpc.RegisterProcedure("hello", p.getHello)
	go p.rpc.Listen()
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

func (p *Peer) printKnownPeers() {
	print(p.Name, " known peers: ")
	for _, addr := range p.KnownPeers() {
		print(addr.String(), " ")
	}
	println()
}
