package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.golang.org")
	if err != nil {
		log.Fatalln("http get error:", err)
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch n.Type {
	case 0:
		fmt.Println("ErrorNode")
	case 1:
		fmt.Println("TextNode")
	case 2:
		fmt.Println("DocumentNode")
	case 3:
		fmt.Println("ElementNode")
	case 4:
		fmt.Println("CommentNode")
	case 5:
		fmt.Println("DoctypeNode")
	case 6:
		fmt.Println("scopeMarkerNode")
	}
	fmt.Println("\t\tType", n.Type)
	fmt.Println("\t\tData", n.Data)
	fmt.Println("\t\tNamespace", n.Namespace)
	fmt.Println("\t\tAttr", n.Attr)
}
