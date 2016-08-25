package evolved

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", index)
}

type Tool struct {
	Name        string
	Description string
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		descrip := req.FormValue("descrip")
		location := req.FormValue("location")

		ctx := appengine.NewContext(req)
		parentKey := datastore.NewKey(ctx, "House", location, 0, nil)
		key := datastore.NewKey(ctx, "Tools", name, 0, parentKey)
		// 1 write / second
		// We will eventually use memcache to buffer our writes

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
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
		<form method="POST">
			<h1>House Location</h1>
			<select name="location">
				<option value="Garage" selected>Garage</option>
				<option value="Kitchen">Kitchen</option>
			</select>
			<h1>Tool</h1>
			<input type="text" name="name"><br>
			<h1>Descrip</h1>
			<textarea name="descrip"></textarea><br>
			<input type="submit">
		</form>`)
}
