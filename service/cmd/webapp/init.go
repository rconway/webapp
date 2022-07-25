package main

import (
	"embed"
	"html/template"
	"io/fs"
	"os"
	"strconv"

	"github.com/rconway/webapp/service/pkg/utils"
)

//================================================================================================================
// Module initialisation
//================================================================================================================

//go:embed www
var wwwFs embed.FS
var wwwRoot, _ = fs.Sub(wwwFs, "www")

var wwwTemplates *template.Template

var servicePort = 8080

func init() {
	// Load html templates
	wwwTemplates = utils.LoadViewTemplates("root", wwwFs, "www/*.html")

	// Listening port number from command-line
	if len(os.Args) > 1 {
		if portNum, err := strconv.Atoi(os.Args[1]); err == nil {
			servicePort = portNum
		}
	}
}
