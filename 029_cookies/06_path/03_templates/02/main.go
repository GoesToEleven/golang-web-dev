package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/dog/bowzer/pictures", bowzerpics)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var c *http.Cookie
	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func bowzer(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:  "user-cookie",
		Value: "this would be the value",
		Path:  "/",
		//Path:  "/dog/bowzer/",
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "bowzer.gohtml", c)
}

func bowzerpics(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "bowzerpics.gohtml", c)
}

func cat(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "cat.gohtml", c)
}
