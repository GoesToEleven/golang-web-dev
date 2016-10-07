package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}

		s := bufio.NewScanner(c)
		for s.Scan() {
			ln := s.Text()
			if ln == "" {
				break
			}
			io.WriteString(c, ln+"\n")
		}

		c.Close()
	}
}
