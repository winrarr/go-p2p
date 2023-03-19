package main

import (
	"p2p/peer"
)

func main() {
	// port := 10000
	// peers := []peer.Peer{peer.CreateGenesisPeer(port)}
	// for i := 1; i < 2; i++ {
	// 	peers = append(peers,
	// 		peer.CreatePeerAndConnect(
	// 			port+i,
	// 			"127.0.0.1", port,
	// 		),
	// 	)
	// }

	// for _, u := range peers[0].KnownPeers() {
	// 	print(u.String(), ", ")
	// }
	// println()

	p1 := peer.CreateGenesisPeer(10000)
	p2 := peer.CreatePeerAndConnect(10001, "12.0.0.1", 10000)

	printKnownPeers(p1)
	printKnownPeers(p2)

	select {}
}

func printKnownPeers(p peer.Peer) {
	for _, addr := range p.KnownPeers() {
		println(addr.String())
	}
}
