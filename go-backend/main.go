package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("react-build")))
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":4444", nil)
	if err != nil { fmt.Println(err) }

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
		<div>
			sign in requested hi
		<div>`,
	); if err != nil { fmt.Println(err) }
}


