package main

import (
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("192.168.1.50"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("Server is running!")

	for {
		buffer := make([]byte, 1024)
		a, b, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		log.Println(a, b)
		log.Println(buffer)
	}
}
