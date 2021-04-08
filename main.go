package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rconway/goutils/httputils"
)

//go:embed app/build
var appRoot embed.FS
var wwwRoot, _ = fs.Sub(appRoot, "app/build")

//================================================================================================================
// Middlewares
//================================================================================================================

func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

//================================================================================================================
// Simulated API endpoint
//================================================================================================================

func newApiRouter(prefix string, router *mux.Router) *mux.Router {
	router.PathPrefix("/fred").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> fred route\n")
	})
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> / route\n")
	})
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> BASE route\n")
	})
	return router
}

//================================================================================================================
// Application (Single Page, e.g. Reactjs)
//================================================================================================================

func newAppRouter(prefix string, router *mux.Router) *mux.Router {
	// Could 'Use' some more middlewares here, if needed

	// Just a single route for the whole app
	router.PathPrefix("").Handler(http.StripPrefix(prefix, httputils.NewSpaHandler(wwwRoot)))
	return router
}

//================================================================================================================
// Entrypoint
//================================================================================================================

func main() {
	router := mux.NewRouter()

	// Register middlewares
	router.Use(loggingMiddleware)

	// API
	newApiRouter("/api", router.PathPrefix("/api").Subrouter())

	// Application (SPA)
	{
		// This is at the root URL path since Reactjs doesn't seem to accommodate a path
		// prefix very easily.
		prefix := ""
		appSubRouter := router.PathPrefix("").Subrouter()
		newAppRouter(prefix, appSubRouter)
	}

	http.ListenAndServe(":8080", router)
}
