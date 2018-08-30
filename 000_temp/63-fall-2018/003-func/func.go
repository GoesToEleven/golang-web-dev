package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	secondStatement()
	finalStat()
}

// func receiver identifier(parameters) returns {code}

func secondStatement() {
	fmt.Println("Here is my second statement")
}

func finalStat() {
	fmt.Println("about to exit")
}
