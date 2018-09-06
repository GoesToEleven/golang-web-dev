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
		license: 7,
		sayings: []string{"Shaken, not stirred", "Bond, James Bond",},
	}
	fmt.Println(p1)

	p2 := person{
		first: "Jenny",
		license: 8,
		sayings: []string{"When Bond can't handle it, call me", "I will always love Bond",},
	}
	fmt.Println(p2)

	xp := []person{p1, p2}
	fmt.Println(xp)

	mp := map[string]person{"Mr":p1, "Ms":p2,}
	fmt.Println(mp)
}