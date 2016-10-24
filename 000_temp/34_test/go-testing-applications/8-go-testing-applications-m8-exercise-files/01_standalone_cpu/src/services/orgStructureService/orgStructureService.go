package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/api/orgUnits", func(w http.ResponseWriter, r *http.Request) {
		var receivers []*Receiver
		if r.URL.Query().Get("type") == "all" {
			receivers = GetReceivers()
		}

		data, _ := json.Marshal(receivers)
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.ListenAndServe(":5000", nil)
}

type Receiver struct {
	Id       int    `json:"id"`
	Location string `json:"location"`
	Type     string `json:"type"`
}

func GetReceivers() []*Receiver {
	result := []*Receiver{
		{Id: 42, Location: "Midville, OH", Type: "Plant"},
		{Id: 627, Location: "Sunnyvale, CA", Type: "Plant"},
		{Id: 10, Location: "Wilmington, VA", Type: "Sales Office"},
	}

	return result
}
