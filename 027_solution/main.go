package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("templates/index.gohtml")
	//	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	/*
		// href="css/reset.css"
		http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("public/css"))))

		// src="pic/surf.jpg"
		http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("public/pic"))))
	*/

	// href="public/css/reset.css"
	// src="public/pic/surf.jpg"
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}
