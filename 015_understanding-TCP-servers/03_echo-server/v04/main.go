package main

import (
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		// handles unlimited connections
		go func() {
			io.Copy(conn, conn)
			conn.Close()
		}()
	}
}
