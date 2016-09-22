package main

import (
	"net"
	"log"
	"io"
	"fmt"
	"bufio"
	"io/ioutil"
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

		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				break
			}
		}

		io.WriteString(c, "hello")
		c.Close()
	}
}
