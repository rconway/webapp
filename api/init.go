package api

import (
	"embed"
	"html/template"
	"log"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

func init() {
	loadViewTemplates()
}

//================================================================================================================
// Load the view templates
//================================================================================================================

//go:embed views
var viewFS embed.FS

var viewTemplates *template.Template

func loadViewTemplates() {
	var err error
	viewTemplates, err = template.ParseFS(viewFS, "views/*")
	if err != nil {
		log.Fatal("ERROR parsing views FS:", err)
	}
}
