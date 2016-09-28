package main

import (
	"net"
	"log"
	"io"
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
			log.Println(err)
		}



		io.WriteString(c, "wassup\n")
		c.Close()
	}
}
