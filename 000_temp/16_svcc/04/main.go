package main

import "fmt"

// composition
// embedded types; inner-type promotion
// methods
// interfaces

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Hello, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	saySomething(p1)

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(sa1)
}
