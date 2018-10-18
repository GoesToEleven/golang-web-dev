package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x", x)
	fmt.Println(&x)

	y := &x

	fmt.Printf("%v\t %T\n", x, x)
	fmt.Printf("%v\t %T\n", y, y)
	fmt.Println(*y)

	*y = 100
	fmt.Println("y", *y)
	fmt.Println("x", x)
}
