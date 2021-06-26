package api

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed swagger-ui
var swaggerRoot embed.FS

//================================================================================================================
// swagger-ui docs
//================================================================================================================

func apiSwaggerHandler(prefix string, router *mux.Router) {
	router.PathPrefix("").Handler(http.StripPrefix(prefix, http.FileServer(http.FS(swaggerRoot))))
}

//================================================================================================================
// Simulated API endpoint
//================================================================================================================

func apiUserHandler(router *mux.Router) {
	router.PathPrefix("/{name}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "zzz: api -> Hello user %v\n", vars["name"])
	})
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> user route\n")
	})
}

func NewApiRouter(prefix string, router *mux.Router) *mux.Router {
	// swagger-ui
	apiSwaggerHandler(prefix, router.PathPrefix("/swagger-ui").Subrouter())
	// /user
	apiUserHandler(router.PathPrefix("/user").Subrouter())
	// /fred
	router.PathPrefix("/fred").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> fred route\n")
	})
	// Root
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewTemplates.ExecuteTemplate(w, "index.html", nil)
	})
	router.Handle("", http.RedirectHandler(prefix+"/", http.StatusPermanentRedirect))
	return router
}
