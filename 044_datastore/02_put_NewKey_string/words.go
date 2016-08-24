package linguist

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", index)
}

type Word struct {
	Term       string
	Definition string
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		term := req.FormValue("term")
		definition := req.FormValue("definition")

		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Word", term, 0, nil)

		entity := &Word{
			Term:       term,
			Definition: definition,
		}

		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<form method="POST">
			<h1>Word</h1>
			<input type="text" name="term"><br>
			<h1>Definition</h1>
			<textarea name="definition"></textarea>
			<input type="submit">
		</form>`)
}
