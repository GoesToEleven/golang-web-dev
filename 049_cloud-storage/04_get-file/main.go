package skyhdd

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"net/http"
)

const gcsBucket = "learning-1130.appspot.com"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/golden", retriever)
}

func handler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	html := `
		<h1>UPLOAD</h1>
	    <form method="POST" enctype="multipart/form-data">
		<input type="file" name="dahui">
		<input type="submit">
	    </form>
	`

	if req.Method == "POST" {

		mpf, hdr, err := req.FormFile("dahui")
		if err != nil {
			log.Errorf(ctx, "ERROR handler req.FormFile: ", err)
			http.Error(res, "We were unable to upload your file\n", http.StatusInternalServerError)
			return
		}
		defer mpf.Close()

		fname, err := uploadFile(req, mpf, hdr)
		if err != nil {
			log.Errorf(ctx, "ERROR handler uploadFile: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		fnames, err := putCookie(res, req, fname)
		if err != nil {
			log.Errorf(ctx, "ERROR handler putCookie: ", err)
			http.Error(res, "We were unable to accept your file\n"+err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		html += `<h1>Files</h1>`
		for k, _ := range fnames {
			html += `<h3><a href="/golden?object=` + k + `">` + k + `</a></h3>`
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, html)
}

func retriever(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	objectName := req.FormValue("object")
	rdr, err := getFile(ctx, objectName)
	if err != nil {
		log.Errorf(ctx, "ERROR retriever getFile: ", err)
		http.Error(res, "We were unable to get the file"+objectName+"\n"+err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer rdr.Close()
	io.Copy(res, rdr)
}
