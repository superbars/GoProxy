package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//Listen requests
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		//Handle connection and request
		go handleConnection(conn)
	}
}

func handleConnection(src net.Conn) {
	//Connect to the target
	dst, err := net.Dial("tcp", "https://whatismyipaddress.com:80")
	if err != nil {
		log.Panic(err)
	}
	defer dst.Close()
	//Copy the source to the host connection
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Panic(err)
		}

	}()
	//Copy response from a target server connection -> source connection
	if _, err := io.Copy(src, dst); err != nil {
		log.Panic(err)
	}
}
