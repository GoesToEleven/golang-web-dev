package pblog

import (
	"encoding/json"
	"net/http"

	"github.com/nu7hatch/gouuid"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

// Session holds the user's data
type Session struct {
	ID       string
	Pictures map[string]string
	req      *http.Request
	res      http.ResponseWriter
	ctx      context.Context
}

func getSession(res http.ResponseWriter, req *http.Request) *Session {

	ctx := appengine.NewContext(req)

	s := new(Session)
	s.Pictures = make(map[string]string)
	s.req = req
	s.res = res
	s.ctx = ctx
	// new
	// https://play.golang.org/p/DKIulO1nOo
	// make
	// http://play.golang.org/p/sjxr-B7wbA
	// https://play.golang.org/p/CsoLHJAPLb
	// https://play.golang.org/p/iUv84nfthy

	cookie, err := req.Cookie("sessionid")
	if err != nil || cookie.Value == "" {
		sessionID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: sessionID.String(),
		}
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {

		// put file names from gcs in s.Pictures
		s.ID = cookie.Value
		s.listBucket()

		// create memcache.Item
		bs, err := json.Marshal(s)
		if err != nil {
			log.Errorf(ctx, "ERROR memcache.Get json.Marshal: %s", err)
		}
		item = &memcache.Item{
			Key:   cookie.Value,
			Value: bs,
		}
	}

	json.Unmarshal(item.Value, s)
	s.ID = cookie.Value

	// store in memcache
	s.putSession()

	return s
}

// method sets
// https://goo.gl/BzkqZ7
func (s *Session) putSession() {
	bs, err := json.Marshal(s)
	if err != nil {
		return
	}

	memcache.Set(s.ctx, &memcache.Item{
		Key:   s.ID,
		Value: bs,
	})

	http.SetCookie(s.res, &http.Cookie{
		Name:  "sessionid",
		Value: s.ID,
	})
}
