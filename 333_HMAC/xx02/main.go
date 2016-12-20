package main

import (
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + `|`,
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.Method == http.MethodPost {
		id := strings.Split(cookie.Value, `|`)[0]
		cookie.Value = id + `|email=` + req.FormValue("email")
	}

	http.SetCookie(res, cookie)

	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <br>
	    `+cookie.Value+`
	  </body>
	</html>`)

}

// NOT GOOD PRACTICE
// adding user data to a cookie
// with no way of knowing whether or not
// they might have altered that data
//
// HMAC would allow us to determine
// whether or not the data in the cookie was altered
//
// however, best to store user data
// on the server
// and keep backups
