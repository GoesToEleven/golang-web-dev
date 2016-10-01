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

		fmt.Println("\nNEW CONNECTION")
		bs := make([]byte, 20)
		fmt.Println("before read", bs)
		conn.Read(bs)
		fmt.Println("after read", bs)
		fmt.Println(string(bs))

		fmt.Println("Code got here.")
		conn.Close()
	}
}
