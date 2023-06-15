package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		// http.StatusNotFound는 404인데 이런 코드는 공식문서를 통해 찾을 수 있다.
		// https://pkg.go.dev/net/http 참고
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	// 파일에 대한 포인터에 부착된 메서드가 Stat(), fileInfo와 err를 반환한다?
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// responseWriter와 request 그리고 파일이름, 파일의 최종 변경시간, 파일자체를 인수로 요청한다.
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
