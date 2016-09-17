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
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
