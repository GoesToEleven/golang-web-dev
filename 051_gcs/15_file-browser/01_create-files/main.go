package skyhdd

import (
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

type demo struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

const gcsBucket = "learning-1130.appspot.com"

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

	d.delFiles()
	d.createFiles()
	d.listFiles()
}

func (d *demo) delFiles() {

	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		if err := d.bucket.Object(obj.Name).Delete(d.ctx); err != nil {
			log.Errorf(d.ctx, "deleteFiles: unable to delete bucket %q, file %q: %v", d.bucket, obj.Name, err)
			return
		}
	}
}

func (d *demo) createFiles() {
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/", "bar/nonce/1", "bar/nonce/2", "bar/nonce/compadre/1", "bar/nonce/compadre/2"} {
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

func (d *demo) listFiles() {

	io.WriteString(d.res, "OBJECTS IN BUCKET\n")

	objs, err := d.bucket.List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.res, obj.Name+"\n")
	}
}
