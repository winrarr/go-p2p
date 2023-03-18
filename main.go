package main

import (
	"p2p/peer"
	"time"
)

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

	peer.CreateGenesisPeer(10000)
	time.Sleep(100 * time.Millisecond)
	peer.CreatePeerAndConnect(10001, "localhost", 10000)

	select {}
}
