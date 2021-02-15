package main

import (
	"fmt"
	"io"
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
		<div>
			this is a sanity check for the backend.
		<div>`,
	); if err != nil { fmt.Println(err) }
}

