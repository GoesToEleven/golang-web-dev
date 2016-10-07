package main

import "fmt"

type hotdog int

func main() {
	x := 42
	var y hotdog
	y = 17
	fmt.Println(x)
	fmt.Println(y)

	y = hotdog(x)
	fmt.Println(x)

	x = int(y)
	fmt.Println(x)
}

/*
func (receiver) identifier(params) (returns) {
	<code>
}
*/
