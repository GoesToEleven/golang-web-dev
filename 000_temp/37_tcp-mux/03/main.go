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
		c, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		defer c.Close()

		bs, err := ioutil.ReadAll(c)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(bs))
	}
}
