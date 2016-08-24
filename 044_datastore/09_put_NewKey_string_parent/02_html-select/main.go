package evolved

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
	}

	res.Header().Set("Content-Type", "text/html")

	if req.Method == "POST" {
		fmt.Fprintln(res, req.FormValue("location"))
	}

	fmt.Fprintln(res, `
		<form method="POST" action="/">
			<select name="location">
				<option value="Garage" selected>Garage</option>
				<option value="Kitchen">Kitchen</option>
			</select>
			<input type="submit">
		</form>`)
}
