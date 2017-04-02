package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var crlf string = "\r\n"

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, crlf +
		"IN-MEMORY DATABASE" +
		crlf + crlf +
		"USE:" +
		crlf +
		"\tSET key value" +
		crlf +
		"\tGET key" +
		crlf +
		"\tDEL key" +
		crlf +
		"EXAMPLE:" +
		crlf +
		"\tSET fav chocolate" +
		crlf +
		"\tGET fav" +
		crlf + crlf + crlf)

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		if (len(ln) < 1) { continue }
		// logic
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s"+crlf, v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE" + crlf)
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND " + fs[0] + crlf)
			continue
		}
	}
}
