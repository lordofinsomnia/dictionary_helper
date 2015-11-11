package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/smartystreets/goconvey/convey"
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
	t.Error(testName, " failed!")
	t.Error("	code: ", response.Code)
	t.Error("	body len: ", len(response.Body.String()))
	t.Error("	body: ", response.Body.String())
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
		fmt.Println("path: " + curRoute.path)
		fmt.Println("body: " + response.Body.String())
		fmt.Println("caption: " + curRoute.caption)
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
		if response.Code != http.StatusOK {
			traceError(curRoute.name, t, response)
		}
	}
}

func TestBdd(t *testing.T) {
	webApp = nil
	Convey("StartServer", t, func() {
		configureServer()
		Convey("Server started?", func() {
			So(webApp, ShouldNotBeNil)
		})
		Convey("Has caption Dictionary Helper", func() {
			response := getResponse(webApp, "/")
			body := response.Body.String()
			So(body, ShouldContainSubstring, "<h2>Dictionary Helper</h2>")
		})
		Convey("Has link sources", func() {
			response := getResponse(webApp, "/")
			body := response.Body.String()
			So(body, ShouldContainSubstring, "sources")
		})
		Convey("Open link sources", func() {
			response := getResponse(webApp, "/sources")
			body := response.Body.String()
			So(body, ShouldContainSubstring, "<h2>Dictionary Helper - Sources</h2>")
		})
	})
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

/*
BenchmarkAllRoutes-8	[martini] Started GET / for
[martini] Completed 200 OK in 158.72µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 140.295µs
[martini] Started GET / for
[martini] Completed 200 OK in 168.926µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 94.797µs
[martini] Started GET / for
[martini] Completed 200 OK in 176.606µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 144.925µs
[martini] Started GET / for
[martini] Completed 200 OK in 163.13µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 140.412µs
[martini] Started GET / for
[martini] Completed 200 OK in 151.014µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 108.611µs
[martini] Started GET / for
[martini] Completed 200 OK in 176.237µs
[martini] Started GET /sources for
[martini] Completed 200 OK in 120.532µs
2000000000	         0.00 ns/op
*/
func BenchmarkAllRoutes(b *testing.B) {
	webApp = nil
	configureServer()
	for _, curRoute := range routes {
		getResponse(webApp, curRoute.path)
	}
}
