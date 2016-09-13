package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := strings.ToLower(scanner.Text())
			bs := []byte(line)
			var rot13 = make([]byte, len(bs))
			for i, v := range bs {
				// ascii 97 - 122
				/// 109 + 13 = 122
				if v <= 109 {
					rot13[i] = v + 13
				} else {
					// 110 + 13 = 123
					//123 - 26 = 97
					rot13[i] = v - 26 + 13
				}
			}
			fmt.Fprintf(conn, "rot13: %s\n", rot13)
		}

	}
}
