package main

import (
	"log"
	"os"
	"text/template"
)

// 템플릿 변수 선언
var tpl *template.Template

// 초기화 함수
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// 배열 선언
func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
