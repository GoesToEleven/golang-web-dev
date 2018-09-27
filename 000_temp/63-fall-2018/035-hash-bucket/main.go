package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	// scan the page
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 2000)
	// Loop over the words
	for scanner.Scan() {
		fmt.Print(scanner.Text(), " - ")
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("***************")
	// for i := 28; i <= 126; i++ {
	// fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	// }
}

func hashBucket(word string) int {
	fmt.Print(word[0], " ---- ")
	fmt.Printf("%#U\n", word[0])
	return int(word[0])
}