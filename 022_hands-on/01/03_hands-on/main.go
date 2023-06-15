package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", h1)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", mclead)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func h1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from a dog!\n")
}

func mclead(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("yong.gohtml")
	if err != nil {
		log.Fatalln("Error parsing", err)
	}
	err = tpl.ExecuteTemplate(w, "yong.gohtml", "mclead")

	if err != nil {
		log.Fatalln("error executing template", err)
	}

	// io.WriteString(w, "Hello from a yongari #2!\n")
}
