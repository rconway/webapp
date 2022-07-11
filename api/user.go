package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//================================================================================================================
// /user route
//================================================================================================================

func apiUserHandler(router *mux.Router) {
	router.PathPrefix("/{name}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "api -> Hello user %v", vars["name"])
	})
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "api -> user route")
	})
}
