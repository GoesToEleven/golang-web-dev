package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./public"))))
}

func foo(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `<img src="assets/grandma.png">`)
}
