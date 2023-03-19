package peer

import (
	"log"
	"net"
	"strings"
)

func (p *Peer) SendSendConnections(addr *net.UDPAddr) {
	p.sendTo("sendConnections", addr)
}

func (p *Peer) getSendConnections(payload ...string) string {
	var str strings.Builder
	str.WriteString("connections ")
	for _, addr := range p.knownPeers {
		str.WriteString(addr.String() + " ")
	}
	return str.String()
}

func (p *Peer) SendHello(addr *net.UDPAddr) {
	p.sendTo("hello", addr)
}

func (p *Peer) getHello(payload ...string) string {
	addr, err := net.ResolveUDPAddr("udp", payload[0])
	if err != nil {
		p.knownPeers = append(p.knownPeers, addr)
	}
	return ""
}

func (p *Peer) getConnections(payload ...string) string {
	for _, addr := range payload {
		UDPAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			log.Fatal("could not resolve address")
		}
		p.knownPeers = append(p.knownPeers, UDPAddr)
	}
	return ""
}
