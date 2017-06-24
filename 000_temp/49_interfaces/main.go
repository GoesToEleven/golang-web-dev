package main

import (
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/000_temp/49_interfaces/species"
	_ "github.com/GoesToEleven/golang-web-dev/000_temp/49_interfaces/species"
)

type person struct {
	first string
}

func (person) Speak() {
	fmt.Println("I'm a person")
}

type secretAgent struct {
	person
	ltk bool
}

func (secretAgent) Speak() {
	fmt.Println("I'm a person and a secret agent")
}

func foo(h species.Human) {
	h.Speak()
}

func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}

	p2 := secretAgent{
		person{
			"James",
		},
		true,
	}

	foo(p1)
	foo(p2)
}
