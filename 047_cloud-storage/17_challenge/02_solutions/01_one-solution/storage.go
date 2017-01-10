package gcs_images

import (
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

//function putFile takes in the fileName(<user>/<object-name>.jpg)
//and slice of bytes(image bytes) then writes it to gcs bucket.
func (a *app) putFile(fileName string, b []byte) {
	writer := a.bucket.Object(fileName).NewWriter(a.ctx)
	//make the file public
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	writer.ContentType = "image/jpg"

	//writes the read bytes to gcs bucket.
	_, err := writer.Write(b)
	if err != nil {
		log.Errorf(a.ctx, "createFile: unable to write data to bucket")
		return
	}
	err = writer.Close()
	if err != nil {
		log.Errorf(a.ctx, "createFile Close")
		return
	}
}
