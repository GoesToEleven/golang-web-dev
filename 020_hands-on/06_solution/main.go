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

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		conn.Close()
	}
}
