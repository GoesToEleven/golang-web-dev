package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "footerThing", nil)
	tpl.ExecuteTemplate(os.Stdout, "headerThing", nil)
}
