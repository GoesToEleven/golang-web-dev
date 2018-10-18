package main

import "fmt"

type hotdog int

var hd hotdog

func main() {
	hd = 42
	fmt.Println(hd)
	fmt.Printf("%T\n", hd)
}
