package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"time"
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

		io.WriteString(c, "HELLO FROM SERVER" + time.Now().String())

		c.Close()
	}
}
