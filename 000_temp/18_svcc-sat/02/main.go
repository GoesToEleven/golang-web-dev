package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `delicately says, "Hello, James."`)
}

func (s secretAgent) speak() {
	fmt.Println(s.fname, s.lname, `says, "Shaken, not stirred."`)
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
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()
}

/*
func (receiver) identifier(params) (returns) {

}
*/
