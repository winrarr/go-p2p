package peer

import (
	"log"
	"net"
	"strings"
)

func (p *Peer) sendConnections(payload ...string) string {
	var str strings.Builder
	for _, addr := range p.knownPeers {
		str.WriteString(addr.String())
	}
	return str.String()
}

func (p *Peer) connections(payload ...string) string {
	for _, addr := range payload {
		UDPAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			log.Fatal("could not resolve address")
		}
		p.knownPeers = append(p.knownPeers, UDPAddr)
	}
	return ""
}
