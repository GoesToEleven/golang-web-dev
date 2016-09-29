package main

import (
	"log"
	"net"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintln(conn, "Hello from tcp server")

		conn.Close()
	}

}
