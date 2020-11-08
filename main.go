package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = pkger.Stat(path)
	if err != nil {
		// If cannot stat file then use the root index file
		r.URL.Path = filepath.Join(h.staticPath, h.indexPath)
	}

	// Use http.FileServer to serve the static dir
	http.FileServer(pkger.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	log.Println("Starting webapp")

	router := mux.NewRouter()

	// API route
	router.PathPrefix("/api").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API path")
	})

	// Static web site
	_ = pkger.Dir("/my-app/build") // force pkger indexing of static files
	spa := spaHandler{staticPath: "/my-app/build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	// Start listening
	port := 8080
	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
