package dictionaryHelper

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepage(t *testing.T) {

	webApp := martini.Classic()
	webApp.Get("/", homeHandler)

	webApp.Use(render.Renderer())

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)

	if response.Body.String() == "" {
		t.Error("TestHomepage failed!")
		t.Error("	code: ", response.Code)
		t.Error("	body: ", response.Body.String())
		t.Error("	body len: ", len(response.Body.String()))
	}
}
