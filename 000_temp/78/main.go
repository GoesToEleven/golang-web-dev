package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
	"io"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}


func main() {
	http.HandleFunc("/", foo)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method", r.Method)

	fn := ""

	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fn = h.Filename

		df, err := os.Create("./assets/"+fn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer df.Close()

		_, err = io.Copy(df, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	 tpl.ExecuteTemplate(w, "index.gohtml", fn)
}
