package main

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/rconway/webapp/service/pkg/utils"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

//go:embed www
var wwwFs embed.FS
var wwwRoot, _ = fs.Sub(wwwFs, "www")

var wwwTemplates *template.Template

func init() {
	wwwTemplates = utils.LoadViewTemplates("root", wwwFs, "www/*.html")
}
