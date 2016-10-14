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

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: "",
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		cookie.Value = email + `|` + getCode(email)
	}

	http.SetCookie(res, cookie)
	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	    `+cookie.Value+`
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	  </body>
	</html>`)

}
