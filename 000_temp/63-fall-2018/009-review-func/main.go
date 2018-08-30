package main

import "fmt"

func main() {
	fmt.Println("This is the entry point beginning of the programming")

	foo()

	// call bar with ARGUMENTS
	a, b := bar("James Bond", 32)
	fmt.Println(a)
	fmt.Println("In 10 years Bond will be", b, "years old.")

	fmt.Println(`Program "about" to exit`)
}

func foo() {
	fmt.Println("Foo is here")
}

// define bar with PARAMETERS
// EXPRESSIONS evaluate to some VALUE, eg, y + 10
// STATEMENT is a line of code to be executed
func bar(x string, y int) (string, int) {
	return fmt.Sprint(x, " is here and he is ", y, " years old"), y + 10
}

// we will pass ARGUMENTS in to a function that has been defined with PARAMETERS

// func receiver identifier(parameters) return(s) {code}
