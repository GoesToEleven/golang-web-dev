package main

import "fmt"

func main() {
	var s interface{}
	s = struct {
		name string
	}{
		name: "james bond",
	}
	fmt.Printf("%T\n", s)
	s = "James Bond"
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
