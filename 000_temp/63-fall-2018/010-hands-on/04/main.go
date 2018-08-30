package main

import "fmt"

func main() {
	z := foo(40, 45)
	fmt.Println(z)
}


// func receiver identifier(parameters) returns {code}
func foo(x int, y int) int {
	return x * y
}