package _1

import "fmt"

type hotdog int
var x hotdog

func main() {
	x = 7
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	y := int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}