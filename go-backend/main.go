package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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
	_, err := io.WriteString(w, formatRequest(r),
	); if err != nil { fmt.Println(err) }
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}


