package main

import "fmt"

func main() {
	a := foo("something")
	fmt.Println(a)

	d := 14
	b := func(z string){
		fmt.Println(z)
	}

	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", b)

	b("james")
	b("jenny")
	x := foo("cat")
	y := foo("bird")
	fmt.Println(x,y)
	fmt.Println(d)
}

func foo(s string) string {
	return s + "dogggggggg"
}