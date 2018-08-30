package main

import "fmt"

func main() {
	foo()
}


// func receiver identifier(parameters) returns {code}
func foo() {
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}