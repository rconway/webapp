package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rconway/goutils/httputils"
	"github.com/rconway/webapp/api"
)

//go:embed app/build
var appFs embed.FS
var appRoot, _ = fs.Sub(appFs, "app/build")

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
		viewTemplates.ExecuteTemplate(w, "index.html", nil)
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
	router.PathPrefix("/").Handler(http.FileServer(http.FS(wwwRoot)))

	// TODO - this is now obsolete
	// Unmatched route - redirect to base path
	router.PathPrefix("").Handler(http.RedirectHandler("/", http.StatusPermanentRedirect))

	http.ListenAndServe(":8080", router)
}
