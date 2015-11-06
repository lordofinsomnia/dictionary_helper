package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startServer() *httptest.ResponseRecorder {
	webApp := martini.Classic()
	webApp.Get("/", homeHandler)

	webApp.Use(render.Renderer())

	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)
	return response
}

func traceError(t *testing.T, response *httptest.ResponseRecorder) {
	t.Error("TestHomepage failed!")
	t.Error("	code: ", response.Code)
	t.Error("	body: ", response.Body.String())
	t.Error("	body len: ", len(response.Body.String()))
}

func TestHomepage(t *testing.T) {
	response := startServer()

	if response.Body.String() == "" {
		traceError(t, response)
	}
}
