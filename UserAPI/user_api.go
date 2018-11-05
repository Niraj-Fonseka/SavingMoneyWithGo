package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8000", nil)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world User!")
}
