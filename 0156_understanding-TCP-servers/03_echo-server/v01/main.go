package main

import "net"

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

		for {
			bs := make([]byte, 1024)
			n, err := conn.Read(bs)
			if err != nil {
				break
			}
			_, err = conn.Write(bs[:n])
			if err != nil {
				break
			}
		}

		conn.Close()
	}
}
