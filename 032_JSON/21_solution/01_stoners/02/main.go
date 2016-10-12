package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type cities struct {
	Data [][]interface{} `json:"data"`
}

func main() {
	nf, err := os.Open("who-is-getting-high.json")
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadAll(nf)
	if err != nil {
		log.Fatalln(err)
	}

	var data cities

	err = json.Unmarshal(bs, &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
