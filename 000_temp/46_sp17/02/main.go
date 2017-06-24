package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")
}
