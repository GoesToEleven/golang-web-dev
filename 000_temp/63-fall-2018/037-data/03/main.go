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

	t := true
	err = tpl.ExecuteTemplate(f1, "default.gohtml", t)
	if err != nil {
		panic(err)
	}


	f := false
	err = tpl.ExecuteTemplate(f2, "about.gohtml", f)
	if err != nil {
		panic(err)
	}
}
