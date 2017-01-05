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
			log.Fatalln(err)
		}

		s := bufio.NewScanner(c)
		for s.Scan() {
			fmt.Println(s.Text())
			break
		}
		io.WriteString(c, "wassup\n")
		c.Close()
	}
}
