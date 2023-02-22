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
		prefix := "/app"

		// The Single-Page-Application should be delivered from the prefix path with a trailing '/'.
		// This is important to ensure that relative paths used within the SPA inherit the full path,
		// including the prefix.
		spaPrefix := prefix + "/"
		appSubRouter := router.PathPrefix(spaPrefix).Subrouter() // Subrouter for SPA
		newAppRouter(prefix, appSubRouter)                       // Handler for SPA

		// For the 'bare' prefix path (with no trailing '/') we redirect to add the trailing '/',
		// being careful to use a relative path here.
		redirectionPrefix := "." + spaPrefix
		router.PathPrefix(prefix).Handler(http.RedirectHandler(redirectionPrefix, http.StatusTemporaryRedirect))
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
