package main

import "fmt"

func main() {
	foo()

	// anonymous func

	func(x int){
		fmt.Println(x)
	}(4)
}

func foo() {fmt.Println("this is foo")}