package main

import "fmt"

func main() {
	xi := []int{1,2,4,5,7,8,9,}
	foo(xi...)
}

func foo(i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}