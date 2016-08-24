//store images from img and photos folder to gcs
//in the form <username>/<object-name>.jpg using
//the below-the-fold webpage.

package gcs_images

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"net/http"
)

const bucket = "learning-1130.appspot.com"

var tpl *template.Template

type app struct {
	ctx      context.Context
	response http.ResponseWriter
	bucket   *storage.BucketHandle
	client   *storage.Client
}

var user_list = [3]string{"user1", "user2", "user3"}

func init() {

	tpl = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", main)

	//serve the css files in a file server instead of uploading it to gcs and querying it.
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
}

func main(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	ctx := appengine.NewContext(request)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "Error in main client err")
	}
	defer client.Close()

	a := &app{
		ctx:      ctx,
		response: response,
		bucket:   client.Bucket(bucket),
		client:   client,
	}

	//upload all the jpgs in the img and photos folder with user prefixes
	a.uploadFiles()

	user_imgs := make(map[string][]string) //map for passing each user's images to template
	for _, user := range user_list {
		query := &storage.Query{
			Prefix: user + "/",
			// TODO: Delimeter field not needed
			//Delimiter: "/",
		}

		objs, err := a.bucket.List(a.ctx, query)
		if err != nil {
			log.Errorf(a.ctx, "ERROR in query_delimiter")
			return
		}

		var s []string
		//store the public links of the queried result to a slice of string
		for _, obj := range objs.Results {
			s = append(s, obj.MediaLink)
		}

		user_imgs[user] = s

	}

	tpl.ExecuteTemplate(response, "index.html", user_imgs)
}
