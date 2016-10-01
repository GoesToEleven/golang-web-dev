package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		err = conn.SetDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			log.Println("CONN TIMEOUT")
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		// now we get here
		// the connection will time out
		fmt.Println("Code got here.")
		conn.Close()
	}
}
