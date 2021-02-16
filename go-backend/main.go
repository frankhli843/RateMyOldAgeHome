package main

import (
	"fmt"
	"net/http"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/services"

	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const PORT string = ":4444"
const DB_URI string = "mongodb+srv://henrik:wvay16pwZsZ0RCvS@cluster0.w708v.mongodb.net/ratemyoldagehome?retryWrites=true&w=majority" 

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

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
