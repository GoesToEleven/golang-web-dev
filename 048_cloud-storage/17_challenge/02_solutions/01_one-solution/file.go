package gcs_images

import (
	"google.golang.org/appengine/log"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//function upload files explicitly passes the jpgs to
//copyFile
func (a *app) uploadFiles() {
	for _, fileName := range []string{
		"img/adventure-lg.jpg",
		"img/adventure-sm.jpg",
		"img/adventure.jpg",
		"photos/0.jpg",
		"photos/2.jpg",
		"photos/3.jpg",
		"photos/4.jpg",
		"photos/5.jpg",
		"photos/6.jpg",
		"photos/7.jpg",
		"photos/8.jpg",
		"photos/10.jpg",
		"photos/11.jpg",
		"photos/12.jpg",
		"photos/13.jpg",
		"photos/14.jpg",
		"photos/15.jpg",
		"photos/16.jpg",
		"photos/17.jpg",
		"photos/18.jpg",
		"photos/19.jpg",
		"photos/20.jpg",
		"photos/21.jpg",
		"photos/22.jpg",
		"photos/23.jpg",
		"photos/24.jpg",
	} {
		a.copyFile(fileName)
	}
}

//function copyFile takes in the string fileName and reads that
//fileName in the system then writes it in the gcs.
//fileName before passing to putFile is formatted to
//<user>/<object-name>.jpg
func (a *app) copyFile(fileName string) {
	//read files from folders img and photos
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf(a.ctx, "Error in copyFile: ", err)
		return
	}

	xs := strings.Split(fileName, "/")
	if xs[0] == "photos" {
		rand.Seed(time.Now().Unix())
		x := rand.Intn(3)
		fileName = user_list[x] + "/" + xs[1]
	}

	//store images to gcs bucket
	a.putFile(fileName, b)
}
