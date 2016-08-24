package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	//	fmt.Println(os.Args[2])
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` +
		name +
		`</h1>
</body>
</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

/*
at the terminal:
go run main.go Todd

or at the terminal:
go build
./03_os-Args Todd

or at the terminal:
go install
03_os-Args Todd
*/
