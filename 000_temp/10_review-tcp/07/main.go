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
			method := words[0]
			uri := words[1]
			if uri == "/" {
				index(c)
			}
			if uri == "/dog" && method == "GET" {
				woof(c)
			}
			if uri == "/dog" && method == "POST" {
				bowwow(c)
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

	body := `
		<!DOCTYPE html>
		<html>
		<body>
		<form action="/dog" method="POST">
		<input name="fname" type="text" placeholder="first name">
		<input name="lname" type="text" placeholder="last name">
		<input type="submit" value="SEND TO SERVER">
		</form>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func bowwow(c net.Conn) {
	fmt.Fprintln(c, "we could process form submission in this func")
}
