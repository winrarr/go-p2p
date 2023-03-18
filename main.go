package main

import "p2p/peer"

func main() {
	// port := 10000
	// peers := []peer.Peer{peer.CreateGenesisPeer("localhost", port)}
	// for i := 1; i < 2; i++ {
	// 	peers = append(peers,
	// 		peer.CreatePeerAndConnect(
	// 			"localhost", port,
	// 			"localhost", port+i,
	// 		),
	// 	)
	// }

	peer.CreateGenesisPeer("localhost", 10000)
	peer.CreatePeerAndConnect("localhost", 10000, "localhost", 10001)

	select {}
}
