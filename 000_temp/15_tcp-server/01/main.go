package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	var n int
	s := bufio.NewScanner(conn)
	for s.Scan() {
		ln := s.Text()

		if n == 0 {
			fmt.Println(ln)
			xs := strings.Fields(ln)
			method := xs[0]
			route := xs[1]
			switch {
			case method == "GET" && route == "/":
				index()
			case method == "GET" && route == "/dog":
				doggy()
			}
		}

		if ln == "" {
			break
		}

		n++
	}

	conn.Close()
}

func index() {
	fmt.Println("YOU ARE AT INDEX")
}

func doggy() {
	fmt.Println("YOU ARE AT DOGGY")
}
