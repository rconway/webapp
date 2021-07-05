package main

import (
	"embed"
	"html/template"

	"github.com/rconway/webapp/utils"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

//go:embed views
var viewFS embed.FS

var viewTemplates *template.Template

func init() {
	viewTemplates = utils.LoadViewTemplates("root", viewFS, "views/*")
}
