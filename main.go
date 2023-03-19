package main

import (
	"p2p/peer"
	"time"
)

func main() {
	_ = peer.CreateGenesisPeer("p1", 10000)
	time.Sleep(100 * time.Millisecond)
	_ = peer.CreatePeerAndConnect("p2", 10001, "127.0.0.1", 10000)
	time.Sleep(100 * time.Millisecond)
	_ = peer.CreatePeerAndConnect("p3", 10002, "127.0.0.1", 10000)
	time.Sleep(100 * time.Millisecond)

	select {}
}

func printKnownPeers(peers ...peer.Peer) {
	for _, peer := range peers {
		print(peer.Name, " known peers: ")
		for _, addr := range peer.KnownPeers() {
			print(addr.String(), " ")
		}
		println()
	}
}
