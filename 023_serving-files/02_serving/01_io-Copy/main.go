package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// 루트에서 dog 함수 호출
	http.HandleFunc("/", dog)
	///toby.jpg url에서 dogPic 호출
	http.HandleFunc("/toby.jpg", dogPic)
	// defaultServeMux를 사용
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	// 헤더에 다음 내용을 써서 리스폰스로 줌
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 다음 이미지 url을 요청함 그러면 위의 HandleFunc 호출
	io.WriteString(w, `
	<img src="/toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	// os에서 파일을 열어서 포인터를 얻는다.
	f, err := os.Open("toby.jpg")
	// 오류처리
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	// 파일 내용을 responsewriter에 담는다.
	io.Copy(w, f)
}
