package handler

import (
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/033_sample-project/templates"
)

// Index handles "/"
func Index(w http.ResponseWriter, req *http.Request) {
	tpl.TPL.ExecuteTemplate()
}
