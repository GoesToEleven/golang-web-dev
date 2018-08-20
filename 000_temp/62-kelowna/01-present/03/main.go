package main

import "fmt"

// https://play.golang.org/p/gLdpEzVm54U

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (returns) {code}

func (p person) speak() {
	fmt.Println(p.first, `says, "Why, James. Helllllo."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.first, `says, "Why Miss Moneypenny. So good to see you. One martini please. Shaken, not stirred."`)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Missy",
		last:  "Moneypenny",
	}

	p2 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(p1)

	fmt.Println(p2)

	p1.speak()

	p2.speak()

	foo(p1)
	foo(p2)
}
