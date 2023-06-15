package main

import (
	"log"
	"net/http"
)

func main() {
	// index.html이 있어야 정적 페이지가 잘 동작한다. inde.html로 바꾸면 전혀 다른 파일서버 형태가 보인다.
	// 0 상태코드 값이 0이면 오류가 아니고 0이 아니면 오류다.
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
