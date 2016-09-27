package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
			log.Println(err)
		}

		// read from connection
		bs, err := ioutil.ReadAll(c)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(bs))

		// WHY IS THIS LINE NEVER WRITTEN?
		io.WriteString(c, "I see you connected")

		c.Close()
	}
}
