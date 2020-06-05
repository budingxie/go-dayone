package main

import (
	"fmt"
	"io"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hi golang")
}

func main() {
	http.HandleFunc("/hello", myWeb)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
