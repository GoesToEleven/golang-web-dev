package memcache

import (
	"encoding/json"
	"github.com/GoesToEleven/golang-web-dev/045_photo-blog/02_solution/03/packages/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func SetPictures(r *http.Request, c *http.Cookie, xs []string) error {
	ctx := appengine.NewContext(r)
	bs, err := json.Marshal(xs)
	errors.Handle(err)
	item1 := memcache.Item{
		Key:   c.Value,
		Value: bs,
	}
	err = memcache.Set(ctx, &item1)
	return err
}

func GetPictures(r *http.Request, c *http.Cookie) ([]string, error) {
	ctx := appengine.NewContext(r)

	item, err := memcache.Get(ctx, c.Value)
	if err != nil {
		return nil, err
	}
	var xs []string
	json.Unmarshal(item.Value, &xs)
	return xs, nil
}
