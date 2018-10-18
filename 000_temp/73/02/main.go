package main

import "fmt"

// AGGREGATE or COMPOSITE
func main() {
	xi := []int{4,5,6,7,42}

	fmt.Println(xi)

	for p, v := range xi {
		fmt.Println(p, v)
	}

	// FOR declaration; condition; post {}
	n := 42
	for {
		fmt.Println(n)
		n++
		if n == 1000 {
			break
		}
	}

	for i := 1000; i <= 2000; i++ {
		fmt.Println(i)
	}
}
