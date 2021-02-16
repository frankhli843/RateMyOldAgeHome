package main

import (
	"fmt"
	"io"
	"net/http"
	"app/util"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("react-build")))
	http.HandleFunc("/api", api)
	err := http.ListenAndServe(":4444", nil)
	if err != nil { fmt.Println(err) }
}

func api(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, util.FormatRequest(r),
	); if err != nil { fmt.Println(err) }
}
