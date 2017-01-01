package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"github.com/GoesToEleven/golang-web-dev/040_mongodb/02_json/models"
)

func main() {
	r := httprouter.New()
	// added route plus parameter
	r.GET("/user/:id", userGet)
	http.ListenAndServe("localhost:8080", r)
}

// changed func name
func userGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}