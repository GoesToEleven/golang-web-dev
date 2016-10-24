package ctrl

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

type HomeController struct{}

func (hc *HomeController) GetHome(w http.ResponseWriter, r *http.Request) {
	page, _ := os.Open("index.html")
	defer page.Close()
	br := bufio.NewReader(page)

	w.Header().Add("Content-Type", "text/html")

	io.Copy(w, br)
}
