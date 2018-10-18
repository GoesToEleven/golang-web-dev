package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a", a)

	foo(a)
	fmt.Println("a after foo", a)

	bar(&a)
	fmt.Println("a after bar", a)
}


func foo(b int) {
	b = 43
	fmt.Println("b", b)
}

func bar(c *int) {
	*c = 44
	fmt.Println("c", c)
	fmt.Println("*c", *c)
}