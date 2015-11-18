package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Route struct {
	name        string
	path        string
	funcHandler gin.HandlerFunc
	caption     string
}

type Link struct {
	path    string
	caption string
	html    string
}

type Page struct {
	Caption string
	Body    string
	Tmpl    *template.Template
}

var links = [...]Link{Link{caption: "Home", path: "/"},
	Link{caption: "Sources", path: "/sources"}}

var routes = [...]Route{
	Route{name: "home",
		path:        "/",
		funcHandler: homeHandler,
		caption:     htmlHeader("Dictionary Helper")},
	Route{name: "source",
		path:        "/sources",
		funcHandler: sourcesHandler,
		caption:     htmlHeader("Dictionary Helper - Sources")}}

var webApp *gin.Engine

func main() {
	configureServer()
	startServer()
}

func startServer() {
	webApp.Run(":3000")
}

func configureServer() {
	webApp = gin.Default()
	webApp.LoadHTMLGlob("templates/*")
	for _, curRoute := range routes {
		webApp.GET(curRoute.path, curRoute.funcHandler)
	}
	/*webApp.Use(render.Renderer(render.Options{
	Directory:  "templates",
	Layout:     "layout",
	Extensions: []string{".tmpl", ".html"},
	Delims:     render.Delims{"{{", "}}"},
	Charset:    "UTF-8",
	IndentJSON: true}))
	*/
}

func homeHandler(c *gin.Context) {
	var homePage Page
	homePage.Caption = ""
	homePage.Body = ""
	homePage.Tmpl = nil
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"Caption": ""})
}
