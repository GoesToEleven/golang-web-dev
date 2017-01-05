package mem

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"
)

type model struct {
	Name     string
	State    bool
	Pictures []string
}

// Model returns a value of type model
func Model(c *http.Cookie, req *http.Request) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	ctx := appengine.NewContext(req)
	bs, err := base64.URLEncoding.DecodeString(usrData)
	if err != nil {
		log.Errorf(ctx, "Error decoding base64: %s", err)
	}

	m := unmarshalModel(bs)

	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != nil {
		m.Pictures = m2.Pictures
		log.Infof(ctx, "PICTURE PATHS RETURNED FROM MEMCACHE %s", m.Pictures)
	}

	return m
}

func marshalModel(m model) []byte {
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func unmarshalModel(bs []byte) model {

	var m model
	err := json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	return m
}
