package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func uploadPhoto(src multipart.File, hdr *multipart.FileHeader, c *http.Cookie) *http.Cookie {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "imgs", fName)
	dst, _ := os.Create(path)
	defer dst.Close()
	src.Seek(0, 0)
	io.Copy(dst, src)
	return addPhoto(fName, c)
}

func addPhoto(fName string, c *http.Cookie) *http.Cookie {
	m := Model(c)
	m.Pictures = append(m.Pictures, fName)

	xs := strings.Split(c.Value, "|")
	id := xs[0]

	cookie := currentVisitor(m, id)
	return cookie
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
