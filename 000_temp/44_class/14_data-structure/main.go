package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) greeting() {
	fmt.Println(p.first, "says hellooooooo James")
}

func (sa secretAgent) greeting() {
	fmt.Println(sa.first, "says shaken, not stirred")
}

type human interface {
	greeting()
}

func sayMore(h human) {
	h.greeting()
}

func main() {
	p1 := person{
		"Eve",
		"Moneypenny",
	}
	p1.greeting()
	sayMore(p1)

	sa1 := secretAgent{
		person{
			"james",
			"bond",
		},
		true,
	}
	sa1.greeting()
	sayMore(sa1)
}
