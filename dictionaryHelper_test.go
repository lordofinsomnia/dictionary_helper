package main

import (
	"github.com/go-martini/martini"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getResponse(webApp *martini.ClassicMartini, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)
	return response
}

func TestUtils(t *testing.T) {
	Convey("Testing utils func", t, func() {
		Convey("htmlHeader func", func() {
			test := "test"
			expected := "<h2>test</h2>"
			So(htmlHeader(test), ShouldEqual, expected)
		})
		Convey("htmlLink func", func() {
			path := "/test"
			caption := "test"
			expected := "<a href=\"/test\">test</a>"
			So(htmlLink(path, caption), ShouldEqual, expected)
		})
		Convey("htmlLabel func", func() {
			name := "test"
			caption := "test"
			expected := "<label for=\"test\">test</label>"
			So(htmlLabel(name, caption), ShouldEqual, expected)
		})
		Convey("htmlInput func", func() {
			name := "test"
			caption := "test"
			expected := "<input type=\"text\">test</input>"
			So(htmlInput(name, caption), ShouldEqual, expected)
		})
	})
}

func TestApp(t *testing.T) {
	webApp = nil
	Convey("StartServer", t, func() {
		configureServer()
		Convey("Server started?", func() {
			So(webApp, ShouldNotBeNil)
		})
		Convey("Routes are set", func() {
			So(routes, ShouldNotBeEmpty)
			for _, curRoute := range routes {
				Convey("Route: "+curRoute.path+" works", func() {
					response := getResponse(webApp, curRoute.path)
					Convey("Http status OK", func() {
						So(response.Code, ShouldEqual, http.StatusOK)
					})
					Convey("Has caption", func() {
						So(response.Body.String(), ShouldContainSubstring, curRoute.caption)
					})
					Convey("Link slice found", func() {
						So(links, ShouldNotBeEmpty)
						Convey("Has all links", func() {
							for _, curLink := range links {
								Convey("Has link: "+curLink.caption, func() {
									link := htmlLink(curLink.path, curLink.caption)
									So(response.Body.String(), ShouldContainSubstring, link)
								})
							}
						})
						Convey("All links works", func() {
							for _, curLink := range links {
								Convey("link ok: "+curLink.caption, func() {
									curLinkResponse := getResponse(webApp, curLink.path)
									So(curLinkResponse.Code, ShouldEqual, http.StatusOK)
								})
							}
						})
					})
				})
			}
		})
	})
}
func TestSources(t *testing.T) {
	webApp = nil
	Convey("StartServer", t, func() {
		configureServer()
		response := getResponse(webApp, "/sources")
		Convey("Sources works", func() {
			Convey("Has all gui items", func() {
				Convey("Has caption", func() {
					Convey("Has caption label", func() {
						label := htmlLabel("caption", "caption")
						So(response.Body.String(), ShouldContainSubstring, label)
					})
					Convey("Has caption editbox", func() {
						editbox := htmlInput("caption", "caption")
						So(response.Body.String(), ShouldContainSubstring, editbox)
					})
				})
				Convey("Has year", func() {
					Convey("Has year label", func() {
						label := htmlLabel("year", "year")
						So(response.Body.String(), ShouldContainSubstring, label)
					})
					Convey("Has year editbox", func() {
						editbox := htmlInput("year", "year")
						So(response.Body.String(), ShouldContainSubstring, editbox)
					})
				})
				Convey("Has shortname", func() {
					Convey("Has shortname label", func() {
						label := htmlLabel("shortName", "shortName")
						So(response.Body.String(), ShouldContainSubstring, label)
					})
					Convey("Has shortname editbox", func() {
						editbox := htmlInput("shortName", "shortName")
						So(response.Body.String(), ShouldContainSubstring, editbox)
					})
				})
			})
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
