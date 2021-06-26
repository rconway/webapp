package api

import (
	"embed"
	"html/template"
	"log"
)

//go:embed views
var viewFS embed.FS

var viewTemplates *template.Template

func init() {
	var err error
	viewTemplates, err = template.ParseFS(viewFS, "views/*")
	if err != nil {
		log.Fatal("ERROR parsing views FS:", err)
	}
}
