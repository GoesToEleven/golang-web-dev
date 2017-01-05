package poms

import (
	"encoding/json"
	"net/http"
	"os"
	"poms/ctrl"
	"poms/model"
	"testing"
)

func TestMain(m *testing.M) {
	ctrl.Setup()

	go http.ListenAndServe(":3000", new(GZipServer))

	m.Run()

	os.Exit(0)
}

func TestGetVendors(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/vendors")

	if err == nil {
		var vendors []model.Vendor
		data := make([]byte, resp.ContentLength)

		resp.Body.Read(data)

		err = json.Unmarshal(data, &vendors)

	}

	if err != nil {
		t.Error("Failed to retrieve vendors")
	}
}
