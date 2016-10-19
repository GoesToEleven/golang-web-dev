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

func (p person) SomeProcessing() int {
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

// Generally speaking, best practice:
// a functions called in a template is formatting only; not processing or data access.

// The main reason you don't want to do any data processing in your template:
// If you're using a function more than once in a template,
// the server needs to do the processing more than once.
// (though the standard library might be cachine processing -
// I've yet to dig into the internals for this -
// if you find the answer, let me know)