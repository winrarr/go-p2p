package main

import "p2p/peer"

func main() {
	port := 10000
	peers := []peer.Peer{peer.CreateGenesisPeer("localhost", port)}
	for i := 0; i < 1; i++ {
		peers = append(peers,
			peer.CreatePeerAndConnect(
				"localhost", port,
				"localholst", port+i,
			),
		)
	}

	peers[1].SendTooAll("sendConnections ")

	select {}
}
