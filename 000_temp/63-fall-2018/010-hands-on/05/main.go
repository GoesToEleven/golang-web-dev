package main

import "fmt"

func main() {
	x := foo("Todd")
	fmt.Println(x)
}


// func receiver identifier(parameters) returns {code}
func foo(s string) string {
	//return fmt.Sprint("Hello ", s, "!")
	return "Hello " + s + "!"
}