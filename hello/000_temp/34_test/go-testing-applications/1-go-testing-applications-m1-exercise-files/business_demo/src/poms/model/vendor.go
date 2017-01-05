package model

import (
	"encoding/json"
	"net/http"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Vendor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Contact *Contact `json:"contact"`
}

func GetVendors() []*Vendor {
	var result []*Vendor

	resp, _ := http.Get("http://localhost:4000/api/vendors?type=manufacturing")

	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)

	json.Unmarshal(data, &result)

	return result
}
