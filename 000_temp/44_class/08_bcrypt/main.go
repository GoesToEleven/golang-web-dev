package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	// user signs up, sending us new credentials
	bs, err := bcrypt.GenerateFromPassword([]byte("shakennotstirred"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// store bs in a database

	//let's pretend they are now logging in
	submittedPassword := "shakennotstirred"
	err = bcrypt.CompareHashAndPassword(bs, []byte(submittedPassword))
	if err != nil {
		log.Println("Passwords do not match:", err)
	}
	fmt.Println("end of program")
}
