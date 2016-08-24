package evolved

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/tools", showTools)
}

type Tool struct {
	Name        string
	Description string
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
	}

	res.Header().Set("Content-Type", "text/html")

	if req.Method == "POST" {
		putTool(res, req)
	}

	fmt.Fprintln(res, `
			<form method="POST" action="/">
				<h1>Tool</h1>
				<input type="text" name="name"><br>
				<h1>Description</h1>
				<textarea name="descrip"></textarea>
				<input type="submit">
			</form>`)
}

func putTool(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	descrip := req.FormValue("descrip")

	ctx := appengine.NewContext(req)
	parentKey := datastore.NewKey(ctx, "House", "Garage", 0, nil)
	key := datastore.NewKey(ctx, "Tools", name, 0, parentKey)
	// 1 write / second

	entity := &Tool{
		Name:        name,
		Description: descrip,
	}

	_, err := datastore.Put(ctx, key, entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func showTools(res http.ResponseWriter, req *http.Request) {

	html := ""
	ctx := appengine.NewContext(req)
	parentKey := datastore.NewKey(ctx, "House", "Garage", 0, nil)
	q := datastore.NewQuery("Tools").Ancestor(parentKey)
	// Ancestor query:
	// Strongly consistent results

	iterator := q.Run(ctx)
	for {
		var entity Tool
		_, err := iterator.Next(&entity)
		if err == datastore.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<dt>` + entity.Name + `</dt>
			<dd>` + entity.Description + `</dd>
		`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `<dl>`+html+`</dl>`)
}
