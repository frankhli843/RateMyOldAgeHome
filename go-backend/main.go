package main

import (
	"fmt"
	"net/http"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/util"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/services"
)

const PORT string = ":4444"

func main() {
	
	collection := util.ConnectDB()

	services.Collection = collection; 

	http.Handle("/", http.FileServer(http.Dir("react-build")))

	// User
	http.HandleFunc("/register", services.Register)
	http.HandleFunc("/login", services.Login)
	http.HandleFunc("/logout", services.Logout)
	http.HandleFunc("/deleteAccount", services.DeleteAccount)

	// Home
	http.HandleFunc("/createHome", services.CreateHome)
	http.HandleFunc("/getHomeList", services.GetHomeList)
	http.HandleFunc("/getHomeByID", services.GetHomeByID)
	http.HandleFunc("/updateHome", services.UpdateHome)
	http.HandleFunc("/deleteHome", services.DeleteHome)

	// Review
	http.HandleFunc("/addReview", services.AddReview)
	http.HandleFunc("/updateReview", services.UpdateReview)
	http.HandleFunc("/removeReview", services.RemoveReview)


	_err := http.ListenAndServe(PORT, nil)
	if _err != nil { fmt.Println(_err) }
}
