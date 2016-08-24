package hello

import (
	"fmt"
	"html/template"
	"net/http"
)

const strForm = `
<html>
  <body>
    <form action="/book" method="POST">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

const tplHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`

var tpl = template.Must(template.New("sign").Parse(tplHTML))

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/book", gbook)
}

func root(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, strForm)
}

func gbook(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, req.FormValue("content"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
