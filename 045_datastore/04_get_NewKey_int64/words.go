package linguist

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"strconv"
	"strings"
)

func init() {
	http.HandleFunc("/", index)
}

var counter int64

type Word struct {
	Term       string
	Definition string
}

func index(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html")

	if req.Method == "POST" {
		putWord(res, req)
	}

	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
		return
	}

	if req.URL.Path != "/" {
		showWord(res, req)
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

	counter++
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Word", "", counter, nil)

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

func showWord(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	var key *datastore.Key
	x := strings.Split(req.URL.Path, "/")[1]
	y, err := strconv.Atoi(x)
	if err != nil {
		key = datastore.NewKey(ctx, "Word", x, 0, nil)
	} else {
		key = datastore.NewKey(ctx, "Word", "", int64(y), nil)
	}

	var entity Word
	err = datastore.Get(ctx, key, &entity)

	if err == datastore.ErrNoSuchEntity {
		http.NotFound(res, req)
		return
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	fmt.Fprintln(res, `
		<dl>
			<dt>`+entity.Term+`</dt>
			<dd>`+entity.Definition+`</dd>
		</dl>
	`)
}
