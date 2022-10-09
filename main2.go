package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	DefaultWriteTimeout = 15 * time.Second
	DefaultReadTimeout  = 15 * time.Second
)

func SessionStore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println(r.Form)
	log.Println(r.PostForm)

	w.Write([]byte("TEST"))
}

func main() {
	port := "3010"
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/session/store", SessionStore).Methods("GET")

	handler := handlers.LoggingHandler(os.Stdout, r)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: DefaultWriteTimeout,
		ReadTimeout:  DefaultReadTimeout,
	}

	log.Println("Starting server on port", port)
	log.Fatal(newServer.ListenAndServe())
}
