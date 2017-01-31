package skyhdd

import (
	"crypto/sha1"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func uploadFile(req *http.Request, mpf multipart.File, hdr *multipart.FileHeader) (string, error) {

	ext, err := fileFilter(req, hdr)
	if err != nil {
		return "", err
	}
	name := getSha(mpf) + `.` + ext

	rand.Seed(time.Now().Unix())
	x := rand.Intn(3)
	fmt.Println(x)

	var pfx string
	switch x {
	case 0:
		pfx = "bob/"
	case 1:
		pfx = "james/"
	case 2:
		pfx = "stacey/"
	}

	name = pfx + name

	mpf.Seek(0, 0)

	ctx := appengine.NewContext(req)
	return name, putFile(ctx, name, mpf)
}

func fileFilter(req *http.Request, hdr *multipart.FileHeader) (string, error) {

	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "FILE EXTENSION: %s", ext)

	switch ext {
	case "jpg", "jpeg", "txt", "md":
		return ext, nil
	}
	return ext, fmt.Errorf("We do not allow files of type %s. We only allow jpg, jpeg, txt, md extensions.", ext)
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
