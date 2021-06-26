package api

import (
	"log"
	"strings"
	"testing"
)

func TestTemplateLoad(t *testing.T) {
	const prefix = "; defined templates are: "
	const expectedNumLoadedTemplates = 1
	loadedTemplatesList := viewTemplates.DefinedTemplates()[len(prefix):]
	loadedTemplates := strings.Split(loadedTemplatesList, ",")
	numLoadedTemplates := len(loadedTemplates)
	if numLoadedTemplates != expectedNumLoadedTemplates {
		log.Println("Loaded templates...")
		for i, name := range loadedTemplates {
			log.Printf("  %0.3v = %v\n", i+1, name)
		}
		t.Errorf("Wrong number of loaded templates: expected %v, got %v\n", expectedNumLoadedTemplates, numLoadedTemplates)
	}
}
