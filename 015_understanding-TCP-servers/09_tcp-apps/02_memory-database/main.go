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
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n"+
		"USE:\n"+
		"SET key value \n"+
		"GET key \n"+
		"DEL key \n\n"+
		"EXAMPLE:\n"+
		"SET fav chocolate \n"+
		"GET fav \n\n\n")

	// read & write
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		data := make(map[string]string)
		// logic
		switch fs[0] {
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
		case "SET":
			if len(fs) != 3 {
				io.WriteString(conn, "EXPECTED VALUE\n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value
		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			io.WriteString(conn, "INVALID COMMAND "+fs[0]+"\n")
		}
	}
}