package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
	"github.com/rconway/goutils/httputils"
)

func main() {
	log.Println("Starting webapp")

	router := mux.NewRouter()

	// API route
	router.PathPrefix("/api").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API path")
	})

	// Static web site
	pkger.Include("/my-app/build") // force pkger indexing of static files
	spa := httputils.SpaHandler{Root: pkger.Dir("/my-app/build"), PathPrefix: "/my-app/build", IndexPath: "index.html", StatFunc: pkger.Stat}
	// spa := httputils.SpaHandler{Root: http.Dir("my-app/build"), PathPrefix: "my-app/build", IndexPath: "index.html", StatFunc: os.Stat}
	router.PathPrefix("/").Handler(spa)

	// Start listening
	port := 8080
	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
