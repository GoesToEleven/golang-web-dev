package main

import "fmt"

func main() {
	var x uint
	x = 18446744073709551615
	fmt.Println(x)
	fmt.Printf("%T\n\n",x)
}
