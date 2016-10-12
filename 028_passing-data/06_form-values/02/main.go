package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			key := "q"
			file, _, err := req.FormFile(key)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			io.Copy(os.Stdout, file)
		}

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		// you have to put this enctype for uploading files
		io.WriteString(res, `<form method="POST" enctype="multipart/form-data">
      <input type="file" name="q">
      <input type="submit">
    </form>`)
	})
	http.ListenAndServe(":8080", nil)
}
