package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/js", js)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func js(w http.ResponseWriter, r *http.Request) {
	nj := struct {
		First string `json:"first-name"`
		Last  string `json:"last-name"`
	}{
		"James",
		"Bond",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(nj)
	if err != nil {
		fmt.Println(err)
		return
	}
	//bs, err := json.Marshal(nj)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//s := string(bs)
	//tpl.ExecuteTemplate(w, "js.gohtml", s)
}
