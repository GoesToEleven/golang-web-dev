package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/GoesToEleven/golang-web-dev/045_photo-blog/02_solution/03/packages/errors"
	"github.com/GoesToEleven/golang-web-dev/045_photo-blog/02_solution/03/packages/memcache"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "IN INDEX") // fyi
	c := GetCookie(w, r)
	xs, err := memcache.GetPictures(r, c)
	errors.Handle(err)

	bs, err := json.Marshal(xs)
	if err != nil {
		log.Infof(ctx, "%v", err)
	}
	log.Infof(ctx, "%s", string(bs)) // fyi
	ctl.tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func (ctl Controller) IndexSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	c := GetCookie(w, r)

	mf, fh, err := r.FormFile("nf")
	errors.Handle(err)
	defer mf.Close()

	// new file
	fname := createSHA(mf, fh)
	nf, err := createNewFile(r, fname)
	errors.Handle(err)
	defer nf.Close()

	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)

	// store filename
	xs, err := memcache.GetPictures(r, c)
	errors.Handle(err)
	xs = append(xs, fname)
	memcache.SetPictures(r, c, xs)

	ctl.tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func createSHA(mf multipart.File, fh *multipart.FileHeader) string {
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	return fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
}

func createNewFile(r *http.Request, name string) (*os.File, error) {
	ctx := appengine.NewContext(r)
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(wd, "public", "pics", name)
	nf, err := os.Create(path)
	log.Infof(ctx, "%s", path) // fyi
	if err != nil {
		return nil, err
	}
	return nf, nil
}
