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

func (p person) shout() {
	fmt.Println(p.fname, p.lname, `says, "OH, YES, JAMES. YES!"`)
}

func (s secretAgent) speak() {
	fmt.Println(s.fname, s.lname, `says, "Shaken, not stirred."`)
}

func (s secretAgent) shout() {
	fmt.Println(s.fname, s.lname, `says, "SHAKEN, NOT STIRRED."`)
}

type human interface {
	speak()
	shout()
}

func saySomething(h human) {
	h.speak()
	h.shout()
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

/*
func (receiver) identifier(params) (returns) {

}
*/
