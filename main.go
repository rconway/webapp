package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting webapp")

	mux := mux.NewRouter()

	// API route
	mux.PathPrefix("/api").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API path")
	})

	// Static web site
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./site")))

	// Start listening
	http.Handle("/", mux)
	port := 8080
	log.Println("Listening on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
