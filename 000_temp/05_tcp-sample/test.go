package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Send data to the server.
	message := "Hello, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	// Receive response from the server.
	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	response := string(buffer[:bytesRead])
	log.Println("Response from server:", response)
}
