package services

import (
	"net/http"
	"app/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if util.IsRequestValid(w, r, "/register") {
		// Logic
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