package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	} else {
		conn.RemoteAddr()
	}
}
