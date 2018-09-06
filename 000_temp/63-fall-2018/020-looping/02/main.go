package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 7, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	for idx := range xi {
		fmt.Println(idx)
	}

	for _, val := range xi {
		fmt.Println("-",val)
	}

	m := map[string]int{"James":32, "Jenny":27,}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for key := range m {
		fmt.Println(key)
	}

	for _, value := range m {
		fmt.Println("-", value)
	}
}