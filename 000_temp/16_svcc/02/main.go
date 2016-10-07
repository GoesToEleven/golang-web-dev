package main

import "fmt"

func main() {

	x := []int{7, 14, 12, 42}
	fmt.Println(x)

	y := map[string]int{"todd": 45, "max": 40}
	fmt.Println(y)
	fmt.Println(y["todd"])

	z := struct {
		fname string
		lname string
	}{
		"Todd",
		"McLeod",
	}
	fmt.Println(z)
	fmt.Println(z.fname)
}
