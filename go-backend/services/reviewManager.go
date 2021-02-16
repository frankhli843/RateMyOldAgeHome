package services

import (
	"net/http"
	"github.com/frankhli843/RateMyOldAgeHome/tree/main/go-backend/util"
)

// Review Route
func AddReview(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/addReview") {
		// Logic
	}
	return;
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/updateReview") {
		// Logic
	}
	return;
}

func RemoveReview(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/removeReview") {
		// Logic
	}
	return;
}