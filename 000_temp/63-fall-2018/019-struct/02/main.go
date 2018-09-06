package main

import "fmt"

type person struct {
	first string
	license int
	sayings []string
}

func main() {
	p1 := person{
		first: "James",
		license: 007,
		sayings: []string{"Shaken, not stirred", "Bond, James Bond",},
	}
	fmt.Println(p1)
}