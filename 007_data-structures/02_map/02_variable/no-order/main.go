package main

import (
	"fmt"
)

func main() {

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad"}

	for k, v := range sages {
		fmt.Println(k, " - ", v)
	}

}
