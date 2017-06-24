package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(p.first)
}

type car struct {
	color string
}

func (c car) speak() {
	fmt.Println("I am color:", c.color)
}

func foo(h human) {
	fmt.Println(h)
}

func main() {
	p1 := person{"McLeod"}
	p2 := person{"Bond"}
	c1 := car{"red"}
	foo(p1)
	foo(p2)
	foo(c1)
}
