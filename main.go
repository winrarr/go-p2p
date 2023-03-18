package main

import (
	"log"
	"net"
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

	go server()
	time.Sleep(time.Second)
	client()
	time.Sleep(time.Second)
}

func server() {
	// create unconnected udp connection on localhost:10000
	lAddr := &net.UDPAddr{
		Port: 10000,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal("could not listen")
	}

	// listen for messages and respond
	for {
		buf := make([]byte, 256)
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("could not read from udp")
		}
		conn.WriteToUDP([]byte("received message"), addr)
	}
}

func client() {
	// create unconnected udp connection on localhost:10001
	lAddr := &net.UDPAddr{
		Port: 10001,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Fatal("could not listen")
	}

	// send message "hello" to server
	rAddr := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 10000,
	}
	n, err := conn.WriteToUDP([]byte("hello"), rAddr)
	println("sent", n, "bytes")
	if err != nil {
		println("could not send message")
		log.Fatal(err)
	}
}
