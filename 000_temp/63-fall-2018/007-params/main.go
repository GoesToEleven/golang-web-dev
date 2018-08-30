package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")

	// pass in an ARGUMENT to the func when you CALL IT
	foo("James Bond")

	x := "Miss Moneypenny"
	// pass in an ARGUMENT to the func when you CALL IT
	foo(x)

	n, _ := fmt.Println("1")

	fmt.Println("number of bytes written", n)
}

// foo takes a VALUE of TYPE string
// we could also say:
// foo has a parameter which is a VALUE of TYPE string
// funcs may be defined with PARAMETER(S)
func foo(name string) {
	fmt.Println("Hello", name)
}
