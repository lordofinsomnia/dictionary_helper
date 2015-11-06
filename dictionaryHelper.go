package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	webApp := martini.Classic()
	webApp.Get("/", homeHandler)
	webApp.Get("/sources", sourcesHandler)
	webApp.Use(render.Renderer())
	webApp.Run()
}

func homeHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}

func sourcesHandler(r render.Render) {
	r.HTML(http.StatusOK, "index", "")
}
