package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	// 특정 경로에 있는 파일 읽어오기
	//http.HandleFunc("/static/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)
}
