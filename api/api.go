package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//================================================================================================================
// API Router
//================================================================================================================

func NewApiRouter(prefix string, router *mux.Router) *mux.Router {

	// swagger-ui
	apiSwaggerHandler(prefix+"/swagger-ui", router.PathPrefix("/swagger-ui").Subrouter())

	// /login
	apiLoginHandler(prefix+"/login", router.PathPrefix("/login").Subrouter())

	// /user
	apiUserHandler(router.PathPrefix("/user").Subrouter())

	// Root
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewTemplates.ExecuteTemplate(w, "index.html", nil)
	})

	// Default - redirect to API root path
	router.PathPrefix("").Handler(http.RedirectHandler(prefix+"/", http.StatusPermanentRedirect))

	return router
}
