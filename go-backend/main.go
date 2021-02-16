package main

import (
	"fmt"
	"net/http"
	"app/services"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("react-build")))

	// Routes
	http.HandleFunc("/register", services.Register)
	http.HandleFunc("/login", services.Login)
	http.HandleFunc("/logout", services.Logout)
	http.HandleFunc("/deleteAccount", services.DeleteAccount)

	err := http.ListenAndServe(":4444", nil)
	if err != nil { fmt.Println(err) }
}
