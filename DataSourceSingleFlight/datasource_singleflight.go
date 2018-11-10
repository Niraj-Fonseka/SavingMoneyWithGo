package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

var cost int64

func main() {
	var requestGroup singleflight.Group

	http.HandleFunc("/fetch_data", RequestGroup{r: &requestGroup}.GetDataHandler)

	log.Println("Serving Port :  8001")
	http.ListenAndServe(":8001", nil)
}

type RequestGroup struct {
	r *singleflight.Group
}

func (rg RequestGroup) GetDataHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)
	value, err, _ := rg.r.Do(r.URL.String(), func() (interface{}, error) {

		go func() {
			time.Sleep(5000 * time.Millisecond)
			rg.r.Forget(r.URL.String())
		}()

		return GetData()
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := value.(string)

	io.WriteString(w, data)

}

func GetData() (string, error) {
	time.Sleep(5 * time.Second)
	cost += 100

	fmt.Printf("Cost : %d \n", cost)
	return fmt.Sprintf("Fetch called : %d", cost), nil
}
