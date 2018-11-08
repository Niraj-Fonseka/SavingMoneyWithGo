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
	http.ListenAndServe(":8000", nil)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := GetData()
	io.WriteString(w, data)

}

func GetData() (string, error) {
	time.Sleep(5 * time.Second)
	cost += 100

	fmt.Printf("Cost : %d \n", cost)
	return fmt.Sprintf("Fetch called : %d", cost), nil
}
