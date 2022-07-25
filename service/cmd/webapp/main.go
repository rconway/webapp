package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rconway/goutils/httputils"
	"github.com/rconway/webapp/service/pkg/api"
)

//go:embed app
var appFs embed.FS
var appRoot, _ = fs.Sub(appFs, "app")

//================================================================================================================
// Middlewares
//================================================================================================================

func loggingMiddleware(h http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

//================================================================================================================
// Application (Single Page, e.g. Reactjs)
//================================================================================================================

func newAppRouter(prefix string, router *mux.Router) *mux.Router {
	// Could 'Use' some more middlewares here, if needed

	// Just a single route for the whole app
	router.PathPrefix("").Handler(http.StripPrefix(prefix, httputils.NewSpaHandler(appRoot)))
	return router
}

//================================================================================================================
// Entrypoint
//================================================================================================================

func main() {
	router := mux.NewRouter()

	// Register middlewares
	router.Use(loggingMiddleware)

	// base path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wwwTemplates.ExecuteTemplate(w, "index.html", nil)
	})

	// API
	api.NewApiRouter("/api", router.PathPrefix("/api").Subrouter())

	// Application (SPA - Reactjs)
	{
		// We can use a prefix for the SPA app, as long as we specify the same prefix in the "homepage" field
		// in the package.json file...
		// ```
		// {
		//   "homepage": "/app"
		// }
		//
		// $ npm run build
		// ```
		prefix := "/app"
		// Create a subrouter at the prefix path
		appSubRouter := router.PathPrefix(prefix).Subrouter()
		// Create the app route handler with the subrouter and the prefix.
		newAppRouter(prefix, appSubRouter)
	}

	// static assets
	defaultContentFunc := func(httpFs http.FileSystem, w http.ResponseWriter, r *http.Request) {
		// By default, redirect to base
		http.Redirect(w, r, "", http.StatusPermanentRedirect)
	}
	router.PathPrefix("/").Handler(httputils.FileServerWithDefault(http.FS(wwwRoot), defaultContentFunc))

	fmt.Printf("Running service at http://0.0.0.0:%v\n", servicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", servicePort), router))
}
