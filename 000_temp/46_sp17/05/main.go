package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo(x)
}

func foo(x int) {
	fmt.Println("hello", x)
}
