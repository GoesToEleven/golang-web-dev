package main

import "fmt"

type person struct {
	fname       string
	lname       string
	secretAgent bool
}

func main() {

	x := []int{7, 14, 12, 42}
	fmt.Println(x)

	y := map[string]int{"todd": 45, "max": 40}
	fmt.Println(y)
	fmt.Println(y["todd"])

	z := person{
		"Todd",
		"McLeod",
		true,
	}
	fmt.Println(z)
	fmt.Println(z.fname)
}
