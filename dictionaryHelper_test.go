package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func startTestServer(path string, pathHandler martini.Handler) *httptest.ResponseRecorder {
	webApp := martini.Classic()
	curPath := path
	webApp.Get(curPath, pathHandler)

	webApp.Use(render.Renderer())

	request, _ := http.NewRequest("GET", curPath, nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)
	return response
}

func traceError(testName string, t *testing.T, response *httptest.ResponseRecorder) {
	if response.Code != http.StatusOK {
		t.Error(testName, " failed!")
		t.Error("	code: ", response.Code)
		t.Error("	body len: ", len(response.Body.String()))
		t.Error("	body: ", response.Body.String())
	}
}

func TestRoutes(t *testing.T) {
	for _, curRoute := range routes {
		response := startServer(curRoute.path, curRoute.funcHandler)
		if response.Code != http.StatusOK {
			traceError(curRoute.name, t, response)
		}
	}
}

func TestHomepageHasCaption(t *testing.T) {
	response := startTestServer("", homeHandler)

	if strings.Contains(response.Body.String(), "Dictionary Helper") == false {
		traceError("TestHomepageHasCaption", t, response)
	}
}

func TestHomeHasSourcesLink(t *testing.T) {
	response := startTestServer("", homeHandler)

	if strings.Contains(response.Body.String(), "Sources") == false {
		traceError("TestHomeHasSourcesLink", t, response)
	}
}

func TestHomeSourcesLinkWorks(t *testing.T) {
	response := startTestServer("sources", sourcesHandler)

	if strings.Contains(response.Body.String(), "Sources") == false {
		traceError("TestHomeHasSourcesLink", t, response)
	}
}

/*
[martini] Completed 200 OK in 216.855µs
[martini] Started GET / for
[martini] Completed 200 OK in 198.589µs
[martini] Started GET / for
[martini] Completed 200 OK in 182.777µs
[martini] Started GET / for
[martini] Completed 200 OK in 164.963µs
[martini] Started GET / for
[martini] Completed 200 OK in 170.234µs
[martini] Started GET / for
[martini] Completed 200 OK in 187.29µs
2000000000	         0.00 ns/op
*/
func BenchmarkHomepage(b *testing.B) {
	startTestServer("", homeHandler)
}
