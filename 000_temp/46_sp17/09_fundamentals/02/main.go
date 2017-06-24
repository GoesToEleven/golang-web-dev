package main

import "fmt"

func main() {
	xi := []int{4, 3, 5, 7, 3, 4, 5}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}

	m := map[string]int{"mcleod": 42, "bond": 32}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
