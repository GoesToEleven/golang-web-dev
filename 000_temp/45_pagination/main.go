package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Env struct {
	Data          []int
	From          int
	To            int
	Amount        int
	ForwardStart  int
	BackwardStart int
}

var tpl *template.Template
var xi []int

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	xi = make([]int, 0, 100)
	for i := 0; i <= 100; i++ {
		xi = append(xi, i)
	}
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// we will show records FROM a certain record, TO a certain record
	next := r.FormValue("next")
	from, err := strconv.Atoi(next)
	if err != nil {
		from = 0
	}

	recordsPerPageToShow := 10
	to := from + recordsPerPageToShow

	// get the data to pass in
	// slice a slice; aka, slicing a slice
	// up to "to" but not including "to" is the way slicing works
	// your slices could be slices of anything: int, string, struct
	var xx []int
	if to == len(xi) {
		// get records to end
		xx = xi[from:]
	} else {
		xx = xi[from:to]
	}

	// get the ForwardStart
	// if we are going "from" "to"
	// then the next group of records would start at "to"
	// but only if there are such records
	var fs int
	if to+recordsPerPageToShow <= len(xi) {
		fs = to
	} else {
		fs = len(xi) - recordsPerPageToShow
	}

	// get the BackwardStart
	// if we are going "from" "to"
	// then the previous group of records would start at "from" minus ten (records shown)
	// but only if there are such records
	var bs int
	if from-recordsPerPageToShow >= 0 {
		bs = from - recordsPerPageToShow
	} else {
		bs = 0
	}

	data := Env{
		Data:          xx,
		From:          from,
		To:            to - 1,
		Amount:        recordsPerPageToShow,
		ForwardStart:  fs,
		BackwardStart: bs,
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}
}
