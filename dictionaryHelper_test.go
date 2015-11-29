package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getResponse(r http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	webApp.ServeHTTP(response, request)
	return response
}

func TestApp(t *testing.T) {
	webApp = nil
	routes = nil
	Convey("StartServer", t, func() {
		configureServer()
		Convey("Server started?", func() {
			So(webApp, ShouldNotBeNil)
		})
		Convey("Routes are set", func() {
			So(routes, ShouldNotBeEmpty)
			Convey("Routes inited", func() {
				homeFound := false
				sourcesFound := false

				for _, curRoute := range routes {
					if curRoute.name == "home" && curRoute.path == "/" {
						homeFound = true
					} else if curRoute.name == "source" && curRoute.path == "/sources" {
						sourcesFound = true
					}
				}
				Convey("Routes home inited", func() {
					So(homeFound, ShouldBeTrue)
				})
				Convey("Routes source inited", func() {
					So(sourcesFound, ShouldBeTrue)
				})
			})
			for _, curRoute := range routes {
				Convey("Route: "+curRoute.path+" works", func() {
					response := getResponse(webApp, "GET", curRoute.path)
					responseStr := response.Body.String()
					Convey("Http status OK", func() {
						So(response.Code, ShouldEqual, http.StatusOK)
					})
					Convey("Link slice found", func() {
						So(links, ShouldNotBeEmpty)
						Convey("Has all links", func() {
							for _, curLink := range links {
								Convey("Has link: "+curLink.caption, func() {
									link := htmlLink(curLink.path, curLink.caption)
									So(responseStr, ShouldContainSubstring, link)
								})
							}
						})
						Convey("All links works", func() {
							for _, curLink := range links {
								Convey("link ok: "+curLink.caption, func() {
									curLinkResponse := getResponse(webApp, "GET", curLink.path)
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
		response := getResponse(webApp, "GET", "/sources")
		responseStr := response.Body.String()
		Convey("Sources works", func() {
			Convey("Has all gui items", func() {
				createSourcePage()
				Convey("Has source groupbox", func() {
					Convey("Has groupbox caption", func() {
						So(responseStr, ShouldContainSubstring, grpSource)
					})
					Convey("Has caption", func() {
						Convey("Has caption label", func() {
							So(responseStr, ShouldContainSubstring, lblCaption)
						})
						Convey("Has caption editbox", func() {
							So(responseStr, ShouldContainSubstring, edtCaption)
						})
					})
					Convey("Has year", func() {
						Convey("Has year label", func() {
							So(responseStr, ShouldContainSubstring, lblYear)
						})
						Convey("Has year editbox", func() {
							So(responseStr, ShouldContainSubstring, edtYear)
						})
					})
					Convey("Has shortname", func() {
						Convey("Has shortname label", func() {
							So(responseStr, ShouldContainSubstring, lblShortName)
						})
						Convey("Has shortname editbox", func() {
							So(responseStr, ShouldContainSubstring, edtShortName)
						})
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
PASS
BenchmarkAllRoutes-8	[GIN] 2015/11/19 - 21:19:07 | 200 |      65.965µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |      13.281µs |  |   GET     /sources
[GIN] 2015/11/19 - 21:19:07 | 200 |      67.487µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |      13.881µs |  |   GET     /sources
[GIN] 2015/11/19 - 21:19:07 | 200 |      64.629µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |      13.672µs |  |   GET     /sources
[GIN] 2015/11/19 - 21:19:07 | 200 |      70.156µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |       14.09µs |  |   GET     /sources
[GIN] 2015/11/19 - 21:19:07 | 200 |      64.254µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |      13.413µs |  |   GET     /sources
[GIN] 2015/11/19 - 21:19:07 | 200 |      70.635µs |  |   GET     /
[GIN] 2015/11/19 - 21:19:07 | 200 |      13.966µs |  |   GET     /sources
2000000000	         0.00 ns/op

*/
func BenchmarkAllRoutes(b *testing.B) {
	webApp = nil
	configureServer()
	gin.SetMode(gin.ReleaseMode)
	for _, curRoute := range routes {
		getResponse(webApp, "GET", curRoute.path)
	}
}
