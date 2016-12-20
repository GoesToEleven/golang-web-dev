package main

import "fmt"

type user struct {
	Password []byte
	First    string
}

var dbUser = map[string]user{}

func main() {
	un := "Bond"
	dbUser[un] = user{[]byte("shaken, not stirred"), "James"}
	u := dbUser[un]
	fmt.Println(u)

	storedPassword := string(u.Password)
	enteredPassword := "shaken, not stirred"

	fmt.Println(storedPassword)
	fmt.Println(enteredPassword)
	fmt.Println(storedPassword == enteredPassword)
}
