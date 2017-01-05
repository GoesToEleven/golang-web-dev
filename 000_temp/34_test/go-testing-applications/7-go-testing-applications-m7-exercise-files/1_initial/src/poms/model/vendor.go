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

func GetVendors() ([]*Vendor, error) {
	var result []*Vendor

	resp, err := http.Get("http://localhost:4000/api/vendors?type=manufacturing")

	if err != nil {
		return nil, err
	}

	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)

	json.Unmarshal(data, &result)

	return result, nil
}
