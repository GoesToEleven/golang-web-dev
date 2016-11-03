package main

import (
	"fmt"
	"os"
)

func main() {
	aa := os.Args
	if len(aa) > 1 {
		fmt.Println(aa[1])
	} else {
		fmt.Println(aa[0])
		fmt.Println("Hello, I am a Gopher")
	}
}
