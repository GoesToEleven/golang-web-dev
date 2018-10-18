package main

import "fmt"


// a VALUE of TYPE

var x int
var s string
var b bool
var i float64

func main() {
	x = 42
	s = "James"
	b = true
	i = 42.123
	fmt.Println(x, s, b, i)

	f := 43
	g := "Bond"
	h := false
	j := 43.342
	fmt.Println(f, g, h, j)

	fmt.Printf("%T\t %T\t %T\t %T\t \n", f, g, h, j)
}