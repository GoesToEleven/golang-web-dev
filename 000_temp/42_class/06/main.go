package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
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
		}

		bs, err := ioutil.ReadAll(conn)
		fmt.Println(string(bs))

		conn.Close()
	}
}
