package main

import "fmt"

type person struct {
	first string
	last  string
	saying string
}

// func (receiver) identifier(parameters) (returns) {code}

func (p person) speak() {
	fmt.Println(p.first, "says", p.saying)
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		saying: "Shaken, not stirred.",
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		saying: "Helllllllo, James.",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("---------")

	fmt.Println(p1.first, "says", p1.saying)
	fmt.Println(p2.first, "says", p2.saying)

	fmt.Println("---------")

	p1.speak()
	p2.speak()
}
