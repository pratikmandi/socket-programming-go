package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to server.")
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		// Send message to server
		conn.Write([]byte(text))

		// Read echoed msg from server
		reply := make([]byte, 1024)
		n, _ := conn.Read(reply)

		fmt.Println("Server echo:", string(reply[:n]))
	}
}
