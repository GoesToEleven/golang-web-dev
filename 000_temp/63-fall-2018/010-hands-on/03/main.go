package main

import "fmt"

func main() {

	// when we call a func, we pass in ARGUMENTS
	foo("Todd", 42)
	foo("James", 32)
	foo("Jenny", 27)
	foo("Cole", 19)
}


// func
// modularize our code and DRY (don't repeat yourself)
// funcs are defined with parameters
// parameters specify a VALUE of a certain TYPE to be passed in when the func is called
func foo(x string, y int) {
	fmt.Println(x, y)
}
