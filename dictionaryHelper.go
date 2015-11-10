package main

import (
	"fmt"
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
	Route{name: "home", path: "/", funcHandler: homeHandler, caption: "<h2>Dictionary Helper</h2>"},
	Route{name: "source", path: "/sources", funcHandler: sourcesHandler, caption: "<h2>Dictionary Helper - Sources</h2>"}}

var webApp *martini.ClassicMartini

func main() {
	configureServer()
	webApp.Run()
}

func configureServer() {
	webApp = martini.Classic()
	fmt.Println("webApp instance")
	for _, curRoute := range routes {
		fmt.Println("webApp adding route name: " + curRoute.name + " path " + curRoute.path)
		webApp.Get(curRoute.path, curRoute.funcHandler)
	}
	fmt.Println("webApp setRender")
	webApp.Use(render.Renderer())

}

func homeHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}

func sourcesHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", " - Sources")
}
