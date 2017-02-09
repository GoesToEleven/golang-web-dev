package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Are you human?")
	fmt.Println("YES or NO")
	var answer string
	_, err := fmt.Scanf("%s", &answer)
	if err != nil {
		panic(err)
	}
	if answer != "YES" {
		fmt.Println("You're not human! What?!?!")
		os.Exit(0)
	}
	fmt.Println("I am glad you are human")
}
