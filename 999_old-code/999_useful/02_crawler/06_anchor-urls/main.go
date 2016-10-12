package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.reddit.com")
	if err != nil {
		log.Fatalln("http get error:", err)
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// recursively go through all of the nodes
	// nn is the next node
	f(n)
}

func f(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		printer(n)
	}
	for nn := n.FirstChild; nn != nil; nn = nn.NextSibling {
		f(nn)
	}
}

func printer(n *html.Node) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			fmt.Println(a.Val)
			break
		}
	}
}
