package peer

import (
	"encoding/json"
	"log"
	"net"
	"p2p/rpc"
)

func (p *Peer) SendHello(addr *net.UDPAddr) {
	p.rpc.SendTo("hello", p.conn.LocalAddr().String(), addr)
}

func (p *Peer) getHello(r *rpc.Request) {
	p.knownPeers = append(p.knownPeers, r.Addr)
}

func (p *Peer) SendSendConnections(addr *net.UDPAddr) {
	p.rpc.SendTo("sendConnections", nil, addr)
}

func (p *Peer) getSendConnections(r *rpc.Request) {
	r.Respond("connections", p.knownPeers)
}

func (p *Peer) getConnections(r *rpc.Request) {
	println(string(r.Payload))
	var addresses []*net.UDPAddr
	err := json.Unmarshal(r.Payload, &addresses)
	if err != nil {
		log.Fatal("error unmarshalling address in getConnections")
	}
	p.knownPeers = append(p.knownPeers, addresses...)
	p.printKnownPeers()
}
