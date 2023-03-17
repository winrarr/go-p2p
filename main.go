package main

import "p2p/peer"

func main() {
	port := 1234
	_ = peer.CreateGenesisPeer("localhost", port)
	p1 := peer.CreatePeerAndConnect("localhost", port)

	p1.SendTooAll("sendConnections tsraiotn sratiort ")

	select {}
}
