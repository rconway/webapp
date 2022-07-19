package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestTemplateLoad(t *testing.T) {
	const prefix = "; defined templates are: "
	const expectedNumLoadedTemplates = 2
	loadedTemplatesList := viewTemplates.DefinedTemplates()[len(prefix):]
	loadedTemplates := strings.Split(loadedTemplatesList, ",")
	for i := range loadedTemplates {
		loadedTemplates[i] = strings.TrimSpace(loadedTemplates[i])
	}
	numLoadedTemplates := len(loadedTemplates)
	if numLoadedTemplates != expectedNumLoadedTemplates {
		log.Println("Loaded templates...")
		for i, name := range loadedTemplates {
			log.Printf("  %0.3v = %v\n", i+1, name)
		}
		t.Errorf("Wrong number of loaded templates: expected %v, got %v\n", expectedNumLoadedTemplates, numLoadedTemplates)
	}
}

func TestUserHandler(t *testing.T) {
	// Input and expected output
	user := "fred"
	expectedResponse := fmt.Sprintf("api -> Hello user %v", user)

	// Make request to handler
	req := httptest.NewRequest(http.MethodGet, "/fred", nil)
	w := httptest.NewRecorder()
	router := mux.NewRouter()
	apiUserHandler(router)
	router.ServeHTTP(w, req)

	// Get response data
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	// Check for expected response
	if string(data) != expectedResponse {
		t.Errorf("expected '%v' got '%v'", expectedResponse, string(data))
	}
}
