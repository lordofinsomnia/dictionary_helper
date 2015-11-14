package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
)

type Route struct {
	name        string
	path        string
	funcHandler martini.Handler
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

var webApp *martini.ClassicMartini

func main() {
	configureServer()
	startServer()
}

func startServer() {
	webApp.Run()
}

func configureServer() {
	webApp = martini.Classic()
	for _, curRoute := range routes {
		webApp.Get(curRoute.path, curRoute.funcHandler)
	}
	webApp.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Delims:     render.Delims{"{{", "}}"},
		Charset:    "UTF-8",
		IndentJSON: true}))

}

func homeHandler(r render.Render) {
	var homePage Page
	homePage.Caption = ""
	homePage.Body = ""
	homePage.Tmpl = nil
	r.HTML(http.StatusOK, "index", homePage)
}

func sourcesHandler(r render.Render) {
	var sourcePage Page
	sourcePage.Caption = packCaption("Sources")
	sourcePage.Body = createSourcePage()
	sourcePage.Tmpl = nil

	r.HTML(http.StatusOK, "index", sourcePage)
}
