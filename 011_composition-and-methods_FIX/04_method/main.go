package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) someProcessing() int {
	return 7
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := person{
			Name: "Ian Fleming",
			Age:  56,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
