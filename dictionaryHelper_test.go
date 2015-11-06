package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
	"strings"
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

func traceError(testName string, t *testing.T, response *httptest.ResponseRecorder) {
	t.Error(testName, " failed!")
	t.Error("	code: ", response.Code)
	t.Error("	body len: ", len(response.Body.String()))
	t.Error("	body: ", response.Body.String())
}

func TestHomepage(t *testing.T) {
	response := startServer()

	if response.Body.String() == "" {
		traceError("TestHomepage", t, response)
	}
}

func TestHomepageHasCaption(t *testing.T) {
	response := startServer()

	if strings.Contains(response.Body.String(), "Dictionary Helper") == false {
		traceError("TestHomepageHasCaption", t, response)
	}
}

/*
[martini] Completed 200 OK in 91.373µs
[martini] Started GET / for
[martini] Completed 200 OK in 69.625µs
[martini] Started GET / for
[martini] Completed 200 OK in 65.437µs
[martini] Started GET / for
[martini] Completed 200 OK in 63.202µs
[martini] Started GET / for
[martini] Completed 200 OK in 72.268µs
[martini] Started GET / for
[martini] Completed 200 OK in 65.616µs
2000000000	         0.00 ns/op
*/
func BenchmarkHomepage(b *testing.B) {
	startServer()
}
