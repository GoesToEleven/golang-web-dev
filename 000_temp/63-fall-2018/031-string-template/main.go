package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"io"
)

func main() {
	// CREATE STRING
	// assign to variable
	name := "James"
	str := `html here` + name + `more html`
	fmt.Println(str)

	// CREATE STRING
	// string print
	s := fmt.Sprint(`mas ` + name + `menos`)
	fmt.Println(s)

	// CREATE FILE
	// io.Copy to the file
	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	io.Copy(nf, strings.NewReader(s))

	// CREATE FILE
	// writestring to file
	nf2, err := os.Create("newfile2.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("whoops2", err)
	}

	fmt.Println("bytes written", n)
}
