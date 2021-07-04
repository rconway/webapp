package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//================================================================================================================
// API Router
//================================================================================================================

func NewApiRouter(prefix string, router *mux.Router) *mux.Router {

	// exact path - base = ""
	router.Handle("", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API - base path")
	}))

	// exact path "/" - redirect to base
	router.Handle("/", http.RedirectHandler(prefix, http.StatusPermanentRedirect))

	//----------------------------------------------------------------------------

	// swagger-ui
	apiSwaggerHandler(prefix+"/swagger-ui", router.PathPrefix("/swagger-ui").Subrouter())

	// /login
	apiLoginHandler(prefix+"/login", router.PathPrefix("/login").Subrouter())

	// /user
	apiUserHandler(router.PathPrefix("/user").Subrouter())

	// unmatched route - redirect to api base
	router.PathPrefix("").Handler(http.RedirectHandler(prefix, http.StatusPermanentRedirect))

	return router
}
