package dstore

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func storeMemc(m model, id string, req *http.Request) error {
	ctx := appengine.NewContext(req)
	bs, err := json.Marshal(m)
	if err != nil {
		log.Errorf(ctx, "ERROR storeMemc json.Marshal: %s", err)
		return err
	}
	item1 := memcache.Item{
		Key:   id,
		Value: bs,
	}
	err = memcache.Set(ctx, &item1)
	if err != nil {
		log.Errorf(ctx, "ERROR storeMemc memcache.Set: %s", err)
		return err
	}

	return nil
}

func retrieveMemc(id string, req *http.Request) (model, error) {

	var m model
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, id)
	if err != nil {
		// get data from datastore
		log.Infof(ctx, "VALUE FROM DATASTORE: %s", err)
		m, err = retrieveDstore(id, req)
		if err != nil {
			return m, err
		}
		// put data in memcache
		storeMemc(m, id, req)
		return m, nil
	}
	// unmarshal from JSON
	log.Infof(ctx, "VALUE FROM MEMCACHE")
	err = json.Unmarshal(item.Value, &m)
	if err != nil {
		log.Errorf(ctx, "ERROR retrieveMemc unmarshal: %s", err)
		return m, err
	}
	return m, nil
}
