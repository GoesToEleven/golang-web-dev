package main

import (
	"net"
	"log"
	"io"
	"fmt"
	"bufio"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}
	defer c.Close()

	io.WriteString(c, "hello from dial in\n")

	s := bufio.NewScanner(c)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
