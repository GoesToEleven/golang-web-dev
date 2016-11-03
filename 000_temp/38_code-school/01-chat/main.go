package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println(os.Args[0])
		fmt.Println("Hello, I am a Gopher")
	}
}
