package main

import (
	"./user"
	"fmt"
	"net/http"
)

var mux = http.NewServeMux()

func showUser(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	userID, isValid := user.GetUser(ctx)
	fmt.Fprintf(res, "User is: %+v, the ID is valid: %t\n", userID, isValid)
}

func main() {
	http.HandleFunc("/", user.GetUserWrapper)
	user.SetHandler(mux)

	mux.HandleFunc("/", showUser)
	http.ListenAndServe(":8080", nil)
}
