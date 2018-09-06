package main

import "fmt"

func main() {
	x := "ABCDE"
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	xr := []rune(x)
	fmt.Println(xr)

	for _, v := range xr {
		fmt.Printf("%d - %b - %#X\n", v, v, v)
	}

}