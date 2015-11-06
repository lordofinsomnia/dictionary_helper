package dictionaryHelper

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
)

func main() {
	webApp := martini.Classic()
	webApp.Get("/", homeHandler)
	webApp.Use(render.Renderer())

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)
}

func homeHandler() string {
	return ""
}
