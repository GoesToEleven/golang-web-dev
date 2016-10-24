package poms

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GZipServer struct{}

func (gs *GZipServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var writer http.ResponseWriter
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gzw := gZipWriter{ResponseWriter: w, WriteCloser: gzip.NewWriter(w)}
		defer gzw.Close()
		writer = gzw
		w.Header().Add("Content-Encoding", "gzip")
	} else {
		writer = w
	}

	http.DefaultServeMux.ServeHTTP(writer, r)
}

type gZipWriter struct {
	http.ResponseWriter
	io.WriteCloser
}

func (gzw gZipWriter) Write(b []byte) (int, error) {
	return gzw.WriteCloser.Write(b)
}
