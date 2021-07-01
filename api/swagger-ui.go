package api

import (
	"embed"
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
