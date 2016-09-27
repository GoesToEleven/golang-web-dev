package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}
	defer c.Close()

	bs, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
