package skyhdd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func putCookie(res http.ResponseWriter, req *http.Request, fname string) (map[string]bool, error) {
	mss := make(map[string]bool)
	cookie, _ := req.Cookie("file-names")
	if cookie != nil {
		bs, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler base64.URLEncoding.DecodeString: %s", err)
		}
		err = json.Unmarshal(bs, &mss)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler json.Unmarshal: %s", err)
		}
	}

	mss[fname] = true
	bs, err := json.Marshal(mss)
	if err != nil {
		return mss, fmt.Errorf("ERROR putCookie json.Marshal: ", err)
	}
	b64 := base64.URLEncoding.EncodeToString(bs)

	// FYI
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "COOKIE JSON: %s", string(bs))

	http.SetCookie(res, &http.Cookie{
		Name:  "file-names",
		Value: b64,
	})
	return mss, nil
}
