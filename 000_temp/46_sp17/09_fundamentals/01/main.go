package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println(p.first)
}

func main() {
	p1 := person{"McLeod"}
	p2 := person{"Bond"}
	p1.speak()
	p2.speak()
}
