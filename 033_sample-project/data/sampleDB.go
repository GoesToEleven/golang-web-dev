package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

var DB = make(map[int]User)

func init() {
	// make a simulated database
	password, err := bcrypt.GenerateFromPassword([]byte("ilovejames"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	DB[1] = User{
		"Miss",
		"Moneypenny",
		"mm",
		password,
	}

	password, err = bcrypt.GenerateFromPassword([]byte("neversaynever"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	DB[2] = User{
		"James",
		"Bond",
		"skyfall",
		password,
	}
}