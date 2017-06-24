package main

import (
	"github.com/GoesToEleven/golang-web-dev/046_mongodb/15_postgres/books"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/books", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)
	http.HandleFunc("/books/json", books.ShowJSON)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
