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

		serve(c)
	}
}

var counter int

func serve(c net.Conn) {
	defer c.Close()
	s := bufio.NewScanner(c)
	i := 1
	for s.Scan() {
		ln := s.Text()
		// we're in the request line
		if i == 1 {
			fmt.Println("HERE'S THE REQUEST LINE:", ln)
			words := strings.Split(ln, " ")
			fmt.Println("HERE'S THE METHOD:", words[0])
			fmt.Println("HERE'S THE URI:", words[1])
			fmt.Println("HERE'S THE HTTP VERSION:", words[2])
		}
		i++
		//fmt.Println(ln)
		//fmt.Fprintln(c, ln)
		if ln == "" {
			break
		}
	}
	io.WriteString(c, "wassup")
	fmt.Fprintln(c, "wassup again")
}
