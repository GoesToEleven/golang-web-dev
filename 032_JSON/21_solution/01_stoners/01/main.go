package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	nf, err := os.Open("who-is-getting-high.json")
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadAll(nf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
