package main

import "fmt"

func main() {
	x := foo()
	y := foo()
	fmt.Printf("%T\t %T\n", x,y)

	func(z string){
		fmt.Println(z)
	}("james")

	x("jenny")
	y("dogggggggggy")

}

func foo() func(string) {
	return func(z string){
		fmt.Println(z)
	}
}