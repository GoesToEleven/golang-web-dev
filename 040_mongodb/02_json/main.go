package main

import (
	"encoding/json"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/040_mongodb/02_json/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	// added route plus parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

// changed func name
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
