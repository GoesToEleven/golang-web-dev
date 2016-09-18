package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		// now handles multiple connections
		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		tx := scanner.Text()
		fmt.Println(tx)
		if tx == "" {
			// when tx is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	io.WriteString(c, "CHECK OUT THE RESPONSE BODY PAYLOAD")
}
