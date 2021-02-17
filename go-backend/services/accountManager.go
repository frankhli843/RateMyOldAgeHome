package services

import (
	"net/http"
	"context"
	"encoding/json"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/util"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/register") {
		w.Header().Set("Content-Type", "application/json")

		var user models.User

		// we decode our body request params
		_ = json.NewDecoder(r.Body).Decode(&user)

		// insert our book model.
		result, err := Collection.InsertOne(context.TODO(), user)

		if err != nil {
			util.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
	return;
}

func Login(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/login") {
		// Logic
	}
	return;
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/logout") {
		// Logic
	}
	return;
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/deleteAccount") {
		// Logic
	}
	return;
}