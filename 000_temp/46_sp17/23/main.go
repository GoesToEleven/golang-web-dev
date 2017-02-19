package main

import (
	"fmt"
	"os"
)

func main() {
	// enter arguments when running program
	// eg, go run main.go todd pizza surfing

	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
