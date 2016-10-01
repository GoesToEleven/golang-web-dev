package main

import (
	"net"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	io.WriteString(conn, "Hi, I am dialing in.")
}