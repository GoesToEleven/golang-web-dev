package main

import "fmt"

var c int
var d = 43

const z = 142

type person struct {
	first string
	age int
	saying string
}

// func receiver identifier(params) returns {code}

func (p person) speak() {
	fmt.Println(p.first, "says", p.saying)
}

type secretagent struct {
	person
	ltk bool
}

func (sa secretagent) speak() {
	fmt.Println(sa.first, "says even more", sa.saying)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}

func main() {
	a := "James"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := fmt.Sprint("Hello", a)
	fmt.Println(b)

	c = 42
	fmt.Println("c:", c)

	d = 43
	fmt.Println("d:", d)
	fmt.Printf("d type: %T\n", d)

	// slice AGGREGATE or COMPOSITE data structure

	ee := []int{2,3,4,7,9,}
	fmt.Println(ee)

	// it was going SEQUENCE
	// now it's going ITERATIVE

	for i, v := range ee {
		fmt.Println(i, v)
		// conditional logic
		if i == 3 {
			fmt.Println("HEY, i is EQUAL TO", i)
		}
	}


	// map AGGREGATE or COMPOSITE data structure
	f := map[string]int{"James":32, "Jenny":27,}
	fmt.Println(f)

	for k, v := range f {
		fmt.Println(k, v)
	}

	// struct AGGREGATE or COMPOSITE data structure
	p1 := person{
		first: "James",
		age: 32,
		saying: "shaken, not stirred",
	}

	p2 := person {
		first: "Jenny",
		age: 27,
		saying: "nobody does it better",
	}

	xp := []person{p1, p2}
	fmt.Println(xp)
	for i2, v2 := range xp {
		fmt.Println(i2, v2)
	}

	// loop init, cond, post
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	sa1 := secretagent{
		person: person{
			first: "Jack",
			age:   29,
			saying: "Blahb blahblabhalbhablhab",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first)
	fmt.Println(sa1.person.first)

	p1.speak()
	p2.speak()
	sa1.speak()
	fmt.Println("--")
	foo(p1)
	foo(p2)
	foo(sa1)

	fmt.Println(z)
}
