package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	webApp := martini.Classic()
	webApp.Get("/", homeHandler)
	webApp.Use(render.Renderer())
	webApp.Run()
}

func homeHandler(r render.Render) {
	r.HTML(200, "index", "")
}
