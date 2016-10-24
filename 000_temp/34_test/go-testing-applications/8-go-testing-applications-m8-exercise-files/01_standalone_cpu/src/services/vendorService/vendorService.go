package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api/vendors", func(w http.ResponseWriter, r *http.Request) {
		var vendors []*Vendor
		if r.URL.Query().Get("type") == "manufacturing" {
			vendors = GetVendors()
		}

		data, _ := json.Marshal(vendors)

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.ListenAndServe(":4000", nil)
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Vendor struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Contact *Contact `json:"contact"`
}

func GetVendors() []*Vendor {
	result := []*Vendor{
		{
			Id:   1,
			Name: "Smith and Jones Mfg",
			Contact: &Contact{
				Name:  "Martha Jones",
				Phone: "+1 (888)555-9639",
			},
		}, {
			Id:   2,
			Name: "Oswald Unlimited",
			Contact: &Contact{
				Name:  "Clara Smith",
				Phone: "(926)555-1798",
			},
		}, {
			Id:   3,
			Name: "Noble Products",
			Contact: &Contact{
				Name:  "Brian Jeffries",
				Phone: "(425)555-1998",
			},
		},
	}
	return result
}
