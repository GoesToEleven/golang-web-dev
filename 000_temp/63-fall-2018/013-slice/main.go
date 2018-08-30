package main

import "fmt"

func main() {
	xi := []int{2,3,4,5,6,7,8}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}

	xs := []string{"James", "Jenny", "M", "Q"}

	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println(i, v)
	}
}
