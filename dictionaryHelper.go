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

func main() {
	startServer()
}

func startServer() {
	webApp := martini.Classic()
	for _, curRoute := range routes {
		webApp.Get(curRoute.path, curRoute.funcHandler)
	}
	webApp.Use(render.Renderer())
	webApp.Run()
}

func homeHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}

func sourcesHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}
