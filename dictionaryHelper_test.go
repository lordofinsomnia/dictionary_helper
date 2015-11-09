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

	return getResponse(webApp, curPath)
}

func getResponse(webApp *martini.ClassicMartini, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", path, nil)
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

func TestRoutesSimulator(t *testing.T) {
	for _, curRoute := range routes {
		response := startTestServer(curRoute.path, curRoute.funcHandler)
		if response.Code != http.StatusOK {
			traceError(curRoute.name, t, response)
		}
	}
}

func TestServerSettings(t *testing.T) {
	webApp = nil
	configureServer()
	if webApp == nil {
		t.Error("TestRoutes failed web=nil")
	}
}

func TestRouteHaveCaption(t *testing.T) {
	webApp = nil
	configureServer()
	for _, curRoute := range routes {
		response := getResponse(webApp, curRoute.path)
		if strings.Contains(response.Body.String(), curRoute.caption) == false {
			traceError(curRoute.name, t, response)
		}
	}
}

func TestRouteLinkWorks(t *testing.T) {

	webApp = nil
	configureServer()
	for _, curRoute := range routes {
		response := getResponse(webApp, curRoute.path)
		if strings.Contains(response.Body.String(), curRoute.caption) == false {
			traceError(curRoute.name, t, response)
		}
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
