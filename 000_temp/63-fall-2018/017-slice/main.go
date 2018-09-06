package main

import "fmt"

func main() {
	x := 42
	y := 43
	xi := []int{x, y}

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", xi)

	fmt.Println(xi)

	a := "James"
	b := "Jenny"
	xs := []string{a, b}

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", xs)
	fmt.Println(xs)

}