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

	d.createFiles()
	d.listFiles()
	d.queryFiles()
}

func (d *demo) queryFiles() {
	io.WriteString(d.res, "\nRETRIEVING QUERY LIMITED FILE NAMES ( MaxResults: 2 each loop )\n")

	// create a query
	q := &storage.Query{
		MaxResults: 2,
	}

	for q != nil {
		fmt.Fprintf(d.res, "\n%s\n", "LOOPING THROUGH")

		objs, err := d.bucket.List(d.ctx, q)
		if err != nil {
			log.Errorf(d.ctx, "%v", err)
			return
		}

		for _, obj := range objs.Results {
			io.WriteString(d.res, obj.Name+"\n")
		}

		// Next is the continuation query to retrieve more
		// results with the same filtering criteria. If there
		// are no more results to retrieve, it is nil.
		q = objs.Next
	}
}

func (d *demo) listFiles() {
	io.WriteString(d.res, "\nRETRIEVING ALL FILE NAMES\n")

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
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/"} {
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
