package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//loggedIn := true
	loggedIn := false

	err := tpl.Execute(os.Stdout, loggedIn)
	if err != nil {
		log.Fatalln(err)
	}
}
