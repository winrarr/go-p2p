package peer

import (
	"fmt"
	"log"
	"net"
	"p2p/rpc"
	"strconv"
)

type Peer struct {
	conns []net.Conn
}

func CreateGenesisPeer(listenIp string, listenPort int) Peer {
	go listen(listenIp, listenPort)

	return Peer{
		conns: []net.Conn{},
	}
}

func CreatePeerAndConnect(ip string, port int, listenIp string, listenPort int) Peer {
	conn, err := net.Dial("udp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("cant connect")
	}

	return Peer{
		conns: []net.Conn{conn},
	}
}

func listen(listenIp string, listenPort int) {
	conn, err := newConnection(listenIp, listenPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rpc := rpc.NewRpc(rpc.DefaultUDPReadWriter(256))
	rpc.RegisterProcedure("sendConnections", sendConnections)
	rpc.Start(conn)
}

func (p *Peer) SendTooAll(message string) {
	for _, conn := range p.conns {
		fmt.Fprintln(conn, message)
	}
}

func newConnection(listenIp string, listenPort int) (*net.UDPConn, error) {
	udpAddr := net.UDPAddr{
		IP:   net.ParseIP(listenIp),
		Port: listenPort,
	}
	return net.ListenUDP("udp", &udpAddr)
}
