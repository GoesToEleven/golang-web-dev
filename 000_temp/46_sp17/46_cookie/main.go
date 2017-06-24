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

func foo(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "wasssup",
		Value:    "tecate",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	io.WriteString(w, "<h1>hello</h1>")
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("wasssup")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	io.WriteString(w, c.Value)
}
