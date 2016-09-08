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

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// YOU CAN INITIALIZE A VARIABLE WITH AN UNDERLYING TYPE OF STRUCT EITHER WAY
	p1 := doubleZero{
		person: person{
			Name: "Ian Fleming",
			Age:  56,
		},
		LicenseToKill: false,
	}

	//p1 := doubleZero{
	//	person{
	//		Name: "Ian Fleming",
	//		Age:  56,
	//	},
	//	false,
	//}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
