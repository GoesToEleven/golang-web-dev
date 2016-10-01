package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	bs := make([]byte, 20)
	conn.Read(bs)
	fmt.Fprintf(conn, "I heard you say: %s", bs)

	fmt.Println("Code got here.")
}