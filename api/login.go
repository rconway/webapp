package api

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//================================================================================================================
// /login route
//================================================================================================================

func apiLoginHandler(prefix string, router *mux.Router) {
	// exact path '/'
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viewTemplates.ExecuteTemplate(w, "login.html", "github/initiate")
	})

	// exact path <blank>
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path.Base(r.URL.Path)+"/", http.StatusPermanentRedirect)
	})

	// unmatched route - NOT FOUND
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, prefix+"/", http.StatusPermanentRedirect)
	})
}
