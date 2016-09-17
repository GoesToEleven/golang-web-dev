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
		fmt.Println(scanner.Text())
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
