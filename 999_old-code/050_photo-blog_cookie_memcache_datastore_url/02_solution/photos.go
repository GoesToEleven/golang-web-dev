package mem

import (
	"crypto/sha1"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io"
	"mime/multipart"
	"net/http"
)

func uploadPhoto(src multipart.File, id string, req *http.Request) error {
	defer src.Close()
	fName := getSha(src) + ".jpg"
	return addPhoto(fName, id, req)
}

func addPhoto(fName string, id string, req *http.Request) error {
	ctx := appengine.NewContext(req)

	// DATASTORE
	md, err := retrieveDstore(id, req)
	if err != nil {
		return err
	}
	md.Pictures = append(md.Pictures, fName)
	err = storeDstore(md, req)
	if err != nil {
		log.Errorf(ctx, "ERROR addPhoto storeDstore: %s", err)
		return err
	}

	// MEMCACHE
	var mc model
	mc, err = retrieveMemc(id, req)
	if err != nil {
		log.Errorf(ctx, "ERROR addPhoto retrieveMemc: %s", err)
		return err
	}
	mc.Pictures = append(mc.Pictures, fName)
	err = storeMemc(mc, req)
	if err != nil {
		log.Errorf(ctx, "ERROR addPhoto storeMemc: %s", err)
		return err
	}

	return nil
}

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}
