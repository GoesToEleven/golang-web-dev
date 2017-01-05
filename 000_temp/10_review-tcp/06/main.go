package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	i := 1
	s := bufio.NewScanner(c)
	for s.Scan() {
		ln := s.Text()
		// we're in the request line
		if i == 1 {
			words := strings.Split(ln, " ")
			// method := words[0]
			uri := words[1]
			switch uri {
			case "/":
				index(c)
			case "/dog":
				woof(c)
			}
		}
		i++
		if ln == "" {
			break
		}
	}
}

func index(c net.Conn) {
	fmt.Fprintln(c, "hello world")
}

func woof(c net.Conn) {
	fmt.Fprintln(c, "wooof woof")
}
