package peer

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type peer struct {
	conns []net.Conn
}

func CreateGenesisPeer(listenIp string, listenPort int) peer {
	go listen(listenIp, listenPort)

	return peer{
		conns: []net.Conn{},
	}
}

func CreatePeerAndConnect(ip string, port int, listenIp string, listenPort int) peer {
	conn, err := net.Dial("udp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("cant connect")
	}

	return peer{
		conns: []net.Conn{conn},
	}
}

func (p *peer) SendTooAll(message string) {
	for _, conn := range p.conns {
		fmt.Fprintln(conn, message)
	}
}

func listen(listenIp string, listenPort int) {
	conn, err := newConnection(listenIp, listenPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	readFromConnection(conn)
}

func readWriter(conn net.UDPConn, respond func([]byte) []byte) {
	buf := make([]byte, 256)
	_, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Fatal()
	}
	conn.WriteToUDP(respond(buf), addr)
}

func readFromConnection(conn *net.UDPConn) {
	buf := make([]byte, 256)
	for {
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		fp := strings.Split(string(buf), " ")
		switch fp[0] {
		case "sendConnections":
			conn.WriteToUDP([]byte("ok"), addr)
		case "ok":
			println("ok")
		}
	}
}

func newConnection(listenIp string, listenPort int) (*net.UDPConn, error) {
	udpAddr := net.UDPAddr{
		IP:   net.ParseIP(listenIp),
		Port: listenPort,
	}
	return net.ListenUDP("udp", &udpAddr)
}
