package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type game struct {
	Score1 int
	Score2 int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	g1 := game{
		Score1: 7,
		Score2: 9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
