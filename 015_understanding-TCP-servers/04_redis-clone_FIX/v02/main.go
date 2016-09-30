package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"fmt"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var x int
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
		if x == 0 {
			fs := strings.Fields(ln)
			switch fs[0] {
			case "GET":
				fmt.Fprintln(conn, "THIS IS THE METHOD", fs[0])
			case "POST":
				fmt.Fprintln(conn, "THIS IS THE METHOD", fs[0])
			default:
				fmt.Fprintln(conn, "NEITHER GET NOR POST")
			}
		}
		x++
	}
}

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
		handle(conn)
	}
}
