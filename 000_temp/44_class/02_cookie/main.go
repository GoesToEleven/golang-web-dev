package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "mcleods go class",
		}
		http.SetCookie(w, c)
	}
	io.WriteString(w, c.String())
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c.Value = "what to say"
	http.SetCookie(w, c)
	io.WriteString(w, c.String())
}
