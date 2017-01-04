package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// precision receives value
//type city struct {
//	Precision                          string `json:"Postal"`
//	Latitude, Longitude                float64
//	Address, City, State, Zip, Country string
//}


// precision doesn't receive value
type city struct {
	Precision string
	Latitude, Longitude float64
	Address, City, State, Zip, Country string
}

type cities []city

func main() {
	var data cities
	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
