package mem

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func newVisitor(req *http.Request) *http.Cookie {
	mm := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String(), req)
}

func currentVisitor(m model, id string, req *http.Request) *http.Cookie {
	mm := marshalModel(m)
	return makeCookie(mm, id, req)
}

func initialModel() []byte {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return marshalModel(m)
}
