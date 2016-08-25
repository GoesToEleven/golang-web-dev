package dstore

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

func storeDstore(m model, id string, req *http.Request) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", id, 0, nil)

	_, err := datastore.Put(ctx, key, &m)
	if err != nil {
		log.Errorf(ctx, "ERROR storeDstore datastore.Put: %s", err)
		return err
	}
	return nil
}

func retrieveDstore(id string, req *http.Request) (model, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Users", id, 0, nil)

	var m model
	err := datastore.Get(ctx, key, &m)
	if err != nil {
		log.Errorf(ctx, "ERROR retrieveDstore datastore.Get: %s", err)
		return m, err
	}
	return m, nil
}
