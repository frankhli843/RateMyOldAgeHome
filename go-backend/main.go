package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var server *http.Server

func main() {
	port := 6543
	readTimeout := 100
	writeTimeout := 100
	idleTimeout := 100

	server = &http.Server{
		Handler:      createRouter(),
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Fatalf("Webserver failed to start on port %d", port)
	}

	//http.Handle("/", http.FileServer(http.Dir("react-build")))
    //http.HandleFunc("/login", login)
	//err := http.ListenAndServe(":6543", nil)
	////err := http.ListenAndServe(":6543", http.FileServer(http.Dir("./react-build")))
	//if err != nil { fmt.Println(err) }

}

func createRouter() *mux.Router {

	router := mux.NewRouter()
	createRouterDispatcher(router)

	//Log unhandled requests for debugging
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugf("route %s not found", r.URL.String())
		_, err := w.Write([]byte("404 page not found"));if err != nil {
			fmt.Println(err)
		}
	})

	return router
}

func createRouterDispatcher(root *mux.Router) {
	dispatcherHost := "/"

	dispatcherRoot := "./react-build"
	dispatcherIndex := "./react-build/index.html"

	s := root.Host(dispatcherHost).Subrouter()

	s.PathPrefix("/").Handler(routes.NewAppServer(dispatcherRoot, dispatcherIndex))
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


