package main

import (
	"net/http"
	"fmt"
	"github.com/alecthomas/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/", bar)
	http.HandleFunc("/", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You are at foo - the request method was: %v", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You are at bar - the request method was: %v", req.Method)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You are at barred - the request method was: %v", req.Method)
}
