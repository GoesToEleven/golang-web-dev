package db

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"github.com/GoesToEleven/golang-web-dev/033_sample-project/packages/model"
)

var DB = make(map[int]model.User)

func init() {
	// make a simulated database
	password, err := bcrypt.GenerateFromPassword([]byte("ilovejames"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	DB[1] = model.User{
		"Miss",
		"Moneypenny",
		"mm",
		password,
	}

	password, err = bcrypt.GenerateFromPassword([]byte("neversaynever"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	DB[2] = model.User{
		"James",
		"Bond",
		"skyfall",
		password,
	}
}