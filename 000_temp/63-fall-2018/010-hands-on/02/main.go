package main

import "fmt"

func main() {
	foo()
}


// func receiver identifier(parameters) return(s) {code}

func foo() {
	x := "James Bond"
	y := 32
	fmt.Println(x, y)
}