package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("react-build")))
	http.HandleFunc("/henrik", henrik)
	err := http.ListenAndServe(":4444", nil)
	if err != nil { fmt.Println(err) }
}

func henrik(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/henrik" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	// And now set a new body, which will simulate the same data we read:
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))


}


