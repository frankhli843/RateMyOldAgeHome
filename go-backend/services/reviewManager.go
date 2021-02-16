package services

import (
	"net/http"
	"app/util"
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