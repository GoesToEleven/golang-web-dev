package main

import "fmt"

var z int = 42

func main() {
	y := 43
	fmt.Println(z, y)
	foo()
}

func foo() {
	fmt.Println(z, z, z, z, z)
}
