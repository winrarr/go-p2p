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

func CreateGenesisPeer(lIp string, lPort int) Peer {
	conn := newConnection(lIp, lPort)
	go listen(conn)

	return Peer{
		conn: conn,
	}
}

func CreatePeerAndConnect(rIp string, rPort int, lIp string, lPort int) Peer {
	conn := newConnection(lIp, lPort)
	go listen(conn)

	rAddr := &net.UDPAddr{
		IP:   net.ParseIP(rIp),
		Port: rPort,
	}
	p := Peer{
		conn:       conn,
		knownPeers: []*net.UDPAddr{rAddr},
	}

	p.AskForConnections(rAddr)

	return p
}

func newConnection(lIp string, lPort int) *net.UDPConn {
	lAddr := &net.UDPAddr{
		IP:   net.ParseIP(lIp),
		Port: lPort,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal()
	}
	return conn
}

func listen(conn *net.UDPConn) {
	rpc := rpc.NewRpc(rpc.DefaultUDPReadWriter(256))
	rpc.RegisterProcedure("sendConnections", sendConnections)
	rpc.Start(conn)
}

func (p *Peer) AskForConnections(addr *net.UDPAddr) {
	p.sendTo("sendConnections ", addr)
}

func (p *Peer) SendTooAll(message string) {
	for _, addr := range p.knownPeers {
		p.sendTo(message, addr)
	}
}

func (p *Peer) sendTo(message string, addr *net.UDPAddr) {
	println("sending \"" + message + "\" from " + p.conn.LocalAddr().String() + " to " + addr.String())
	p.conn.WriteToUDP([]byte(message), addr)
}

func (p *Peer) Address() net.Addr {
	return p.conn.LocalAddr()
}
