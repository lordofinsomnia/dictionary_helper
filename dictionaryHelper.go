package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
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

func htmlHeader(withOutHeader string) string {
	return "<h2>" + withOutHeader + "</h2>"
}

func htmlLink(path string, caption string) string {
	return "<a href=\"" + path + "\">" + caption + "</a>"
}

func htmlLabel(name string, caption string) string {
	return "<label for=\"" + name + "\">" + caption + "</label>"
}

func htmlGroupBox(caption string, controls string) string {
	return "<fieldset><legend>" + caption + "</legend>" + controls + "</fieldset>"
}

func htmlInput(name string, caption string) string {
	return "<input type=\"text\">" + caption + "</input>"
}

func htmlAddNewLine(html string) string {
	return html + "<br>"
}

func configureServer() {
	webApp = martini.Classic()
	for _, curRoute := range routes {
		webApp.Get(curRoute.path, curRoute.funcHandler)
	}
	webApp.Use(render.Renderer())
}

func homeHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}

func sourcesHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", " - Sources")
}
