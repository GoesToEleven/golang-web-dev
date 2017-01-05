package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() string {
	return "Yo, what's up"
}

func (sa secretAgent) speak() string {
	return "shaken, not stirred"
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	saySomething(sa1)
}
