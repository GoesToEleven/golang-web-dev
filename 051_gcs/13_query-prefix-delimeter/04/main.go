package skyhdd

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const gcsBucket = "learning-1130.appspot.com"

type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

func handler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer client.Close()

	d := &demo{
		ctx:    ctx,
		res:    res,
		client: client,
		bucket: client.Bucket(gcsBucket),
	}

	d.createFiles()
	d.listFiles()
	io.WriteString(d.res, "\nOBJECT NAMES WITH DELIMITER QUERY ( Delimeter: / )\n")
	d.listDelim()

}

func (d *demo) listDelim() {

	query := &storage.Query{
		Delimiter: "/",
	}

	objs, err := d.bucket.List(d.ctx, query)
	if err != nil {
		log.Errorf(d.ctx, "listBucketDirMode: unable to list bucket %q: %v", gcsBucket, err)
		return
	}

	for _, obj := range objs.Results {
		fmt.Fprintf(d.res, "%v\n", obj.Name)
	}

	fmt.Fprintf(d.res, "\nEVERY OTHER OBJECT HAS A DELIMETER IN FRONT OF IT \nIT'S AS IF EVERY OTHER OBJECT IS IN A FOLDER \nAND YOU ARE ONLY LOOKING AT THE OBJECTS IN THIS FOLDER \nTO SEE THE OTHER OBJECTS, YOU WILL NEED TO ADD A PREFIX QUERY\n\nPREFIXES FOUND WITH THIS QUERY ( storage.ObjectList Prefixes )\n%v", objs.Prefixes)
}

func (d *demo) listFiles() {
	io.WriteString(d.res, "ALL OBJECT NAMES\n")

	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.res, obj.Name+"\n")
	}
}

func (d *demo) createFiles() {
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/", "foo/boo/foo3", "foo/boo/foo/4", "boo/yah5", "compadre/amigo/diaz6", "compadre/luego/hasta7", "bar/nonce/8", "bar/nonce/9", "bar/nonce/compadre/10", "bar/nonce/compadre/11"} {
		d.createFile(n)
	}
}

func (d *demo) createFile(fileName string) {

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Errorf(d.ctx, "createFile: unable to close bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
}
