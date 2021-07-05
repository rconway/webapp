package api

import (
	"embed"
	"html/template"

	"github.com/rconway/webapp/utils"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

func init() {
	viewTemplates = utils.LoadViewTemplates("api", viewFS, "views/*")
}

//================================================================================================================
// Load the view templates
//================================================================================================================

//go:embed views
var viewFS embed.FS

var viewTemplates *template.Template
