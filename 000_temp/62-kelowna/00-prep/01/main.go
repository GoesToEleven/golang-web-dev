package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("hello from", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("hello from", sa.first)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func bar(h human) {
	h.speak()
}

func main() {
	x := person{
		first: "miss",
		last:  "money",
	}

	y := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	x.speak()
	y.speak()

	foo(x)
	foo(y)

	fmt.Println("Hello", x)
	fmt.Println("Hello", y)
}
