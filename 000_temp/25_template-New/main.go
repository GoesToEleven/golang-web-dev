package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.New("somet").Parse("here is the text in the template"))
	tpl.ExecuteTemplate(os.Stdout, "somet", nil)
}
