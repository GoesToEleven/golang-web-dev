package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	i := 0
	var url string

	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			method := strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
			fmt.Println("\nREQUEST LINE:", ln)
			fmt.Println("METHOD:", method)
			fmt.Println("URL:", url)
		} else {
			break
		}
		i++
	}

	body := url
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(conn)
	}
}
