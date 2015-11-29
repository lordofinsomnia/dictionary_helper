package main

import (
	//	"fmt"
	//	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	//	"path/filepath"
)

type Route struct {
	name        string
	path        string
	funcHandler gin.HandlerFunc
}

type Link struct {
	path    string
	caption string
	html    string
}

type Page struct {
	Caption string
	Body    string
	Tmpl    string
}

var links = [...]Link{Link{caption: "Home", path: "/"},
	Link{caption: "Sources", path: "/sources"}}

var routes []Route

var webApp *gin.Engine
var indexTempl *template.Template
var sourceTempl *template.Template
var templates map[string]*template.Template

func main() {
	configureServer()
	startServer()
}

func startServer() {
	webApp.Run(":3000")
}

func initRoutes() {
	routes = make([]Route, 2)
	routeHome := Route{name: "home", path: "/", funcHandler: homeHandler}
	routeSoures := Route{name: "source", path: "/sources", funcHandler: sourcesHandler}
	routes[0] = routeHome
	routes[1] = routeSoures
}

func configureServer() {
	templates = nil
	webApp = gin.Default()
	webApp.LoadHTMLGlob("templates/*.tmpl")
	initRoutes()
	for _, curRoute := range routes {
		webApp.GET(curRoute.path, curRoute.funcHandler)
	}
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{})
}
