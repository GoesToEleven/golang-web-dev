package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hs", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		file, _, err := req.FormFile(key)
		if err != nil {
			log.Println("Error with file", err)
			return
		}
		defer file.Close()
		bs, _ := ioutil.ReadAll(file)
		fmt.Println(string(bs))
		http.Redirect(res, req, "/", 303)
	})

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		// you have to put this enctype for uploading files
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data"
		action="/hs">
      <input type="file" name="q">
      <input type="submit">
    </form>`)
	})
	http.ListenAndServe(":8080", nil)
}
