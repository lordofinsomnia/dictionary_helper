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

/*
// Load templates on program initialisation
func initTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := "./templates"

	fmt.Println("templates:")
	fmt.Println(templatesDir)
	templateFiles, err := filepath.Glob(templatesDir + "/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, templateFile := range templateFiles {
		templates[filepath.Base(templateFile)] = template.Must(template.ParseFiles(templateFiles...))
	}
	fmt.Println("templates:")
	fmt.Println(templates)
	fmt.Println("template files:")
	fmt.Println(templateFiles)

}

func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "base", data)

	return nil
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	for curTemplateName, _ := range templates {
		r.Add(curTemplateName, templates[curTemplateName])
	}
	fmt.Println("templates:")
	fmt.Println(templates)
	fmt.Println("r:")
	fmt.Println(r)
	return r
}*/

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{})
}
