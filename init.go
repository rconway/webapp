package main

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/rconway/webapp/utils"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

//go:embed www
var wwwFs embed.FS
var wwwRoot, _ = fs.Sub(wwwFs, "www")

var viewTemplates *template.Template

func init() {
	viewTemplates = utils.LoadViewTemplates("root", wwwFs, "www/*.html")
}
