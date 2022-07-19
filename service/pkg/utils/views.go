package utils

import (
	"html/template"
	"io/fs"
	"log"
)

//================================================================================================================
// Helper to load the view templates
//================================================================================================================

func LoadViewTemplates(owner string, fs fs.FS, pattern string) *template.Template {
	tmpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		log.Fatalf("[%v] ERROR parsing views FS: %v\n", owner, err)
	}
	return tmpl
}
