package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	var code string

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		code = email + `|` + getCode(email)
	}

	io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <form method="POST">
    `+code+`
      <input type="email" name="email">
      <input type="submit">
    </form>
  </body>
</html>`)

}
