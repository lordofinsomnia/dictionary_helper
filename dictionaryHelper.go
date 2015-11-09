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
}

var routes = [...]Route{
	Route{name: "home", path: "/", funcHandler: homeHandler},
	Route{name: "source", path: "/sources", funcHandler: sourcesHandler}}

var webApp *martini.ClassicMartini

func main() {
	configureServer()
	webApp.Run()
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
	r.HTML(http.StatusOK, "index", "")
}
