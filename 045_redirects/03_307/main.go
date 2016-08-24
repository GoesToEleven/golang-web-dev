package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("form.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", canine)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}

	if req.Method == "POST" {
		const url = "/dog"
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
	}

	err := tpl.ExecuteTemplate(res, "form.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func canine(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "What's up dog.")
}
