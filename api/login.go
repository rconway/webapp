package api

import (
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//================================================================================================================
// /login route
//================================================================================================================

func apiLoginHandler(prefix string, router *mux.Router) {
	// exact path - base = ""
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		githubInitiateUrl := path.Base(r.URL.Path) + "/github/initiate"
		redirectUrl := r.URL.Query().Get("redirect_url")
		if len(redirectUrl) > 0 {
			log.Println("Got redirect_url =", redirectUrl)
			githubInitiateUrl += "?redirect_url=" + redirectUrl
		} else {
			log.Println("NO redirect_url =>", r.URL.Query())
		}
		viewTemplates.ExecuteTemplate(w, "login.html", githubInitiateUrl)
	})

	// exact path "/" - redirect to base
	router.Handle("/", http.RedirectHandler(prefix+"", http.StatusPermanentRedirect))

	// unmatched route - redirect to login base
	router.PathPrefix("").Handler(http.RedirectHandler(prefix, http.StatusPermanentRedirect))
}
