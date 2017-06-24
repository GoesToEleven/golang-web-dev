package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

type customer struct {
	Name string `json:"name" bson:"name"`
}

var tpl *template.Template
var DB *mgo.Database
var Customer *mgo.Collection

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	// get a mongo sessions
	//s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Customer = DB.C("customer")

	fmt.Println("You connected to your mongo database.")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	customers := []customer{}
	err := Customer.Find(bson.M{}).All(&customers)
	if err != nil {
		fmt.Println(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", customers)
}
