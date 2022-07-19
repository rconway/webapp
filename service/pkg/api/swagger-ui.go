package api

import (
	"embed"
	"io/fs"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//go:embed swagger-ui
var swaggerFs embed.FS
var swaggerRoot, _ = fs.Sub(swaggerFs, "swagger-ui")

//================================================================================================================
// swagger-ui docs
//================================================================================================================

func apiSwaggerHandler(prefix string, router *mux.Router) {
	router.PathPrefix("/").Handler(http.StripPrefix(prefix, http.FileServer(http.FS(swaggerRoot))))

	// exact path <blank> - redirect to '/'
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path.Base(r.URL.Path)+"/", http.StatusPermanentRedirect)
	})

	// unmatched route - redirect to '/'
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, prefix+"/", http.StatusPermanentRedirect)
	})
}
