package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var cost int64

func main() {
	http.HandleFunc("/fetch_data", GetDataHandler)
	log.Println("Serving Port :  8000")
	http.ListenAndServe(":8000", nil)
}
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	data, _ := GetData()
	io.WriteString(w, data)
}

func GetData() (string, error) {
	time.Sleep(5 * time.Second)
	cost += 100
	fmt.Printf("Cost : %d \n", cost)
	return fmt.Sprintf("Fetch called : %d", cost), nil
}
