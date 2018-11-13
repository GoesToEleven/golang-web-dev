package main

import "fmt"

func main() {
	xi := []int{4, 5, 6, 7}

	for n, i := range xi {
		fmt.Println(n, i)
	}

	m := map[string]int{
		"mcleod": 47,
		"bond":   27,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
