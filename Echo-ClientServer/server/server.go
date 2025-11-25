package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to listen")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Couldn't accept connection")
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		// Read message from client
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		fmt.Printf("Received: %s", msg)

		// Echo message back
		conn.Write([]byte(msg))
	}
}
