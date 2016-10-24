package main

import (
	"net/http"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/031_sessions/04_chaining_FIX/mcookie"
)

func Auth1(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		fmt.Println("here is the authorization code")
		c := mcookie.GetCookie(req)
		cl := db[c.Value]
		if !cl.Loggedin {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		h(w, req)
	}
}

// hmmm
// https://play.golang.org/p/I_BDJusVYv