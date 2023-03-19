package peer

import (
	"encoding/json"
	"log"
	"net"
	"p2p/rpc"
)

func (p *Peer) SendHello(addr *net.UDPAddr) {
	p.rpc.SendTo("hello", p.conn.LocalAddr(), addr)
}

func (p *Peer) getHello(r *rpc.Request) {
	var addr *net.UDPAddr
	err := json.Unmarshal(r.Payload, addr)
	if err != nil {
		log.Fatal("error unmarshalling address in getHello")
	}
	p.knownPeers = append(p.knownPeers, addr)
}

func (p *Peer) SendSendConnections(addr *net.UDPAddr) {
	p.rpc.SendTo("sendConnections", nil, addr)
}

func (p *Peer) getSendConnections(r *rpc.Request) {
	r.Respond("connections", p.knownPeers)
}

func (p *Peer) getConnections(r *rpc.Request) {
	var connections []*net.UDPAddr
	err := json.Unmarshal(r.Payload, &connections)
	if err != nil {
		log.Fatal("error unmarshalling connections in getConnections")
	}
	p.knownPeers = append(p.knownPeers, connections...)
}
