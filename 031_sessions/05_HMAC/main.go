package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	code := getCode("test@example.com")
	fmt.Println(code)
}
