package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
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
	    <a href="/authenticate">Validate This HMAC</a>
	  </body>
	</html>`)

}

func auth(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		http.Redirect(res, req, "/", 303)
		return
	}

	if cookie.Value == "" {
		http.Redirect(res, req, "/", 303)
		return
	}

	xs := strings.Split(cookie.Value, "|")
	fmt.Println("HERE'S THE SLICE", xs)
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)
	//codeCheck := getCode(email + "s")

	if codeRcvd != codeCheck {
		log.Println("HMAC codes didn't match")
		log.Println(codeRcvd)
		log.Println(codeCheck)
		http.Redirect(res, req, "/", 303)
		return
	}

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheck+` - RECALCULATED </h1>
	  </body>
	</html>`)
}
