package main

import "fmt"

// slice
// map
// struct

type person struct {
	fname string
	lname string
}

func main() {
	xi := []int{2, 3, 4, 5, 6, 7}
	xi = append(xi)
	fmt.Println(xi)

	xs := make([]string, 0, 100)
	xs = append(xs, "Go rocks.")
	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))

	for i, v := range xi {
		fmt.Println(i, v)
	}

	y := map[string]string{
		"Todd": "chocolate",
		"Jim":  "vanilla",
	}
	y["Mark"] = "chocolate also"

	fmt.Println(y)

	for k, v := range y {
		fmt.Println(k, v)
	}

	z := person{
		"James",
		"Bond",
	}
	fmt.Println(z)
	fmt.Println(z.fname)
	fmt.Printf("%T\n", z)
}
