package services

import (
	"net/http"
	"app/util"
)

// Home Routes
func CreateHome(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/createHome") {
		// Logic
	}
	return;
}

func GetHomeList(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/getHomeList") {
		// Logic
	}
	return;
}

func GetHomeByID(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/getHomeByID") {
		// Logic
	}
	return;
}

func UpdateHome(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/updateHome") {
		// Logic
	}
	return;
}

func DeleteHome(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/deleteHome") {
		// Logic
	}
	return;
}
