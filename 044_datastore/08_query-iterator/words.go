package linguist

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/words", showWords)
}

type Word struct {
	Term       string
	Definition string
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
		return
	}

	res.Header().Set("Content-Type", "text/html")

	if req.Method == "POST" {
		putWord(res, req)
	}

	fmt.Fprintln(res, `
			<form method="POST" action="/">
				<h1>Word</h1>
				<input type="text" name="term"><br>
				<h1>Definition</h1>
				<textarea name="definition"></textarea>
				<input type="submit">
			</form>`)
}

func putWord(res http.ResponseWriter, req *http.Request) {
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

func showWords(res http.ResponseWriter, req *http.Request) {

	var html string
	q := datastore.NewQuery("Word").Limit(3).Order("-Term")
	ctx := appengine.NewContext(req)

	iterator := q.Run(ctx)
	for {
		var entity Word
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		html += `
			<dt>` + entity.Term + `</dt>
			<dd>` + entity.Definition + `</dd>
		`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `<dl>`+html+`</dl>`)
}
