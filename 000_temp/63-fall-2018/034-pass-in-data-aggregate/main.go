package main

import (
	"text/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	First string
	Last string
	Age int
}

func main() {

	p1 := person{
		First: "James",
		Last: "Bond",
		Age: 32,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", p1)
	if err != nil {
		panic(err)
	}

	p2 := person{
		First: "Jenny",
		Last: "Moneypenny",
		Age: 27,
	}

	xp := []person{p1, p2}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", xp)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "lasagne.php", p1)
	if err != nil {
		panic(err)
	}

}
