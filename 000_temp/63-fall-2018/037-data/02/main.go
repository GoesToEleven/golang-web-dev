package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	f1, err := os.Create("default.html")
	if err != nil {
		log.Fatal("f1 errored", err)
	}
	defer f1.Close()

	f2, err := os.Create("about.html")
	if err != nil {
		log.Fatal("f1 errored", err)
	}
	defer f2.Close()

	s := "      here is text with space     "
	err = tpl.ExecuteTemplate(f1, "default.gohtml", s)
	if err != nil {
		panic(err)
	}


	t := `    here is a raw 

string 


literal      


`
	err = tpl.ExecuteTemplate(f2, "about.gohtml", t)
	if err != nil {
		panic(err)
	}
}
