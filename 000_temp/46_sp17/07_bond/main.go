package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() string {
	return fmt.Sprintln("uptown, func you up, uptown func you up")
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln("shaken, not stirred")
}

type human interface {
	speak() string
}

func foo(h human) {
	fmt.Println("hello from foo")
}

func main() {
	p1 := person{"Nina", "Simone", 25}
	sa1 := secretAgent{person{"Ian", "Fleming", 42}, false}
	foo(p1)
	foo(sa1)
}
