package main

import "fmt"

func main() {
	m := map[string]int{"James":32, "Jenny":27,}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
