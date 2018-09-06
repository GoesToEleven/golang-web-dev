package main

import "fmt"

func main() {
	m := map[string]int{"James":7, "Jenny":8,}
	fmt.Printf("%T\n", m)
	fmt.Println(m)

	fmt.Println(m["James"])
}