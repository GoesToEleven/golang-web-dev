package main

import "fmt"

type person struct {
	fname    string
	lname    string
	age      int
	greeting string
}

func (p person) speak() {
	fmt.Printf("%s %s says, %s\n", p.fname, p.lname, p.greeting)
}

func main() {
	p1 := person{"Miss", "Moneypenny", 24, "Good morning, James. It's soooo good to see you."}
	p2 := person{"Olivia", "Mansfield", 45, "Bond, what the hell are you doing?"}
	p1.speak()
	p2.speak()
}
