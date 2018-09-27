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
	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", nil)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "about.gohtml", nil)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	f, err := os.Create("main.html")
	if err != nil {
		log.Fatal("another error", err)
	}
	defer f.Close()

	f2, err := os.Create("about.html")
	if err != nil {
		log.Fatal("another error", err)
	}
	defer f2.Close()

	err = tpl.ExecuteTemplate(f, "main.gohtml", nil)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	err = tpl.ExecuteTemplate(f2, "about.gohtml", nil)
	if err != nil {
		log.Fatal("there was an error", err)
	}


}