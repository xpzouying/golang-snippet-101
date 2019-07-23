package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/hi", handleFunc)

	// 8080 为端口号
	fmt.Println(http.ListenAndServe(":8080", nil))
}
