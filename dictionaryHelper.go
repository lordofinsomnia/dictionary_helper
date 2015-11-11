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
