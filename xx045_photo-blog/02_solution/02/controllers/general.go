package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Controller struct {
	tpl *template.Template
}

func NewController(tpl *template.Template) *Controller {
	return &Controller{tpl}
}

func (ctl Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := getCookie(w, r)
	xs := strings.Split(c.Value, "|")
	ctl.tpl.ExecuteTemplate(w, "index.gohtml", xs[1:]) //only send over images
}

func (ctl Controller) IndexSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := getCookie(w, r)

	mf, fh, err := r.FormFile("nf")
	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()

	// create a new file
	fname := createSHA(mf, fh)
	nf, err := createNewFile(fname)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()

	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)

	// add filename to this user's cookie
	c = appendValue(w, c, fname)

	//only send over images
	xs := strings.Split(c.Value, "|")
	ctl.tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func createSHA(mf multipart.File, fh *multipart.FileHeader) string {
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	return fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
}

func createNewFile(name string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	path := filepath.Join(wd, "public", "pics", name)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	return nf, nil
}
