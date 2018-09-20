package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foooooooo")
}

func bar() {
	fmt.Println("barrrrrrrrr")
	defer one()
	two()
}

func one() {
	fmt.Println("oneeeeeee")
}

func two() {
	fmt.Println("twoooooooooo")
}