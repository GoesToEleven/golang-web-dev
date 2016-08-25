package pblog

import (
	"crypto/sha1"
	"fmt"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
)

func (s *Session) uploadPhoto() {

	// retrieve the submitted file
	mpf, hdr, err := s.req.FormFile("data")
	if err != nil {
		log.Errorf(s.ctx, "ERROR uploadPhoto s.req.FormFile: %s", err)
		http.Redirect(s.res, s.req, `/?id=`+s.ID, http.StatusSeeOther)
		return
	}
	defer mpf.Close()

	// only allow jpeg or jpg
	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]
	log.Infof(s.ctx, "FILE EXTENSION: %s", ext)
	switch ext {
	case "jpg", "jpeg":
		log.Infof(s.ctx, "GOOD FILE EXTENSION: %s", ext)
	default:
		log.Errorf(s.ctx, "We do not allow files of type %s. We only allow jpg, jpeg extensions.", ext)
		http.Redirect(s.res, s.req, `/?id=`+s.ID, http.StatusSeeOther)
		return
	}

	// make a file name
	h := sha1.New()
	io.Copy(h, mpf)
	name := fmt.Sprintf("%x", h.Sum(nil)) + `.` + ext
	log.Infof(s.ctx, "FILE NAME: %s", name)
	mpf.Seek(0, 0)

	// put file
	client, err := storage.NewClient(s.ctx)
	if err != nil {
		log.Errorf(s.ctx, "ERROR uploadPhoto storage.NewClient: %s", err)
		http.Redirect(s.res, s.req, `/?id=`+s.ID, http.StatusSeeOther)
		return
	}
	defer client.Close()
	writer := client.Bucket(gcsBucket).Object(name).NewWriter(s.ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, mpf)
	err = writer.Close()
	if err != nil {
		log.Errorf(s.ctx, "ERROR uploadPhoto writer.Close: %s", err)
		http.Redirect(s.res, s.req, `/?id=`+s.ID, http.StatusSeeOther)
		return
	}

	// update session
	s.Pictures[name] = name
	s.putSession()
}

func (s *Session) listBucket() {

	client, err := storage.NewClient(s.ctx)
	if err != nil {
		log.Errorf(s.ctx, "ERROR listBucket storage.NewClient: %v", err)
		return
	}
	defer client.Close()

	q := &storage.Query{
		Prefix: s.ID,
	}

	objs, err := client.Bucket(gcsBucket).List(s.ctx, q)
	if err != nil {
		log.Errorf(s.ctx, "ERROR listBucket client.Bucket: %v", err)
		return
	}

	for _, obj := range objs.Results {
		s.Pictures[obj.Name] = obj.Name
	}

	s.putSession()
}
