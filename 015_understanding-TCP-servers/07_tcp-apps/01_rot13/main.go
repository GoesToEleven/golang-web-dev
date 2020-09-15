package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"unicode"
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
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		r := rot13(ln)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(ln string) string {
	var r13 = make([]rune, len(ln))
	for i, v := range ln {
		if unicode.IsLower(v) && v < unicode.MaxASCII {
			r13[i] = rotateCharacter(v, 'a')
		} else if unicode.IsUpper(v) && v < unicode.MaxASCII {
			r13[i] = rotateCharacter(v, 'A')
		} else {
			r13[i] = v
		}
	}
	return string(r13)
}

func rotateCharacter(v rune, base rune) rune {
	return (v - base + 13) % 26 + base
}
