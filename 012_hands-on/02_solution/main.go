package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	years := []year{
		{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"CSCI-40", "Introduction to Programming in Go", "4"},
					{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"CSCI-50", "Advanced Go", "5"},
					{"CSCI-190", "Advanced Web Programming with Go", "5"},
					{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatalln(err)
	}
}
