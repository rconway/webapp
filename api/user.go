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
		fmt.Fprintf(w, "zzz: api -> Hello user %v\n", vars["name"])
	})
	router.PathPrefix("").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zzz: api -> user route\n")
	})
}
