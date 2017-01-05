package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		bs, err := ioutil.ReadAll(c)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(bs))

		io.WriteString(c, "hello")
		c.Close()
	}
}
