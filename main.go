package main

import (
	"p2p/peer"
)

func main() {
	port := 10000
	peers := []peer.Peer{peer.CreateGenesisPeer(port)}
	for i := 1; i < 2; i++ {
		peers = append(peers,
			peer.CreatePeerAndConnect(
				port+i,
				"127.0.0.1", port,
			),
		)
	}

	for _, u := range peers[0].KnownPeers() {
		print(u.String(), ", ")
	}
	println()

	select {}
}
