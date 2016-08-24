package mem

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func uploadPhoto(src multipart.File, hdr *multipart.FileHeader, c *http.Cookie, req *http.Request) *http.Cookie {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "assets", "imgs", fName)
	dst, _ := os.Create(path)
	defer dst.Close()
	src.Seek(0, 0)
	io.Copy(dst, src)
	return addPhoto(fName, c, req)
}

func addPhoto(fName string, c *http.Cookie, req *http.Request) *http.Cookie {
	xs := strings.Split(c.Value, "|")

	// memcache
	id := xs[0]
	m2 := retrieveMemc(req, id)
	m2.Pictures = append(m2.Pictures, fName)
	mm := marshalModel(m2)
	b64 := base64.URLEncoding.EncodeToString(mm)
	storeMemc([]byte(b64), id, req)

	// cookie
	usrData := xs[1]
	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		ctx := appengine.NewContext(req)
		log.Errorf(ctx, "Error decoding base64: %s", err)
	}
	m := unmarshalModel(bs)
	m.Pictures = append(m.Pictures, fName)
	mm = marshalModel(m)
	b64 = base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64) // hmac
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
