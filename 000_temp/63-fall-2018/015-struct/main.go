package main

import "fmt"

type person struct {
	first string
	last string
	age int
	sayings []string
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
		sayings: []string{"Shaken, not stirred", "Bond, James Bond",},
	}
	fmt.Println(p1)

	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 27,
		sayings: []string{"Danger knows no gender", "A woman's place is in control",},
	}
	fmt.Println(p2)

	xp := []person{p1, p2}

	fmt.Println("---------")
	for i, v := range xp {
		fmt.Println(i, v, v.first)
		for j, w := range v.sayings {
			fmt.Println(j, w)
		}
	}
	fmt.Println("---------")

	m := map[string]person{"James":p1, "Jenny":p2,}

	for k, p := range m {
		fmt.Println(k, p)
	}
}
