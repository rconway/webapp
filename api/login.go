package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//================================================================================================================
// /login route
//================================================================================================================

func apiLoginHandler(router *mux.Router) {
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewTemplates.ExecuteTemplate(w, "login.html", nil)
	})
}
