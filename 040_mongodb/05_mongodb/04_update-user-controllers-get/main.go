package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/GoesToEleven/golang-web-dev/040_mongodb/05_mongodb/04_update-user-controllers-get/controllers"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:80", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}