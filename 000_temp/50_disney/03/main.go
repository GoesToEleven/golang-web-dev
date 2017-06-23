package main

import "fmt"

var x int = 42

func init() {
	fmt.Println("in init ", x)
}

func main() {
	fmt.Println("in main ", x)
}
