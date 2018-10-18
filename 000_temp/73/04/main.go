package main

import "fmt"

type person struct {
	first string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p := person{
		first: "Miss Moneypenny",
		age: 27,
	}

	sa := secretAgent {
		person: person {
			first: "James",
			age: 32,
		},
		ltk: true,
	}

	fmt.Println(p)
	fmt.Println(sa)
}