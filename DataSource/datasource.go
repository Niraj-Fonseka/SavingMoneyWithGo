package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var cost int64

func main() {

	http.HandleFunc("/fetch_data", helloWorldHandler)
	http.ListenAndServe(":8001", nil)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	data := GetData()
	io.WriteString(w, data)

}

func GetData() string {
	time.Sleep(5 * time.Second)
	cost += 100
	return fmt.Sprintf("Fetch called : %d", cost)
}
