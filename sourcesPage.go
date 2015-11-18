package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/martini-contrib/render"
	"net/http"
)

var lblCaption string
var edtCaption string
var lblYear string
var edtYear string
var lblShortName string
var edtShortName string
var controls string
var grpSource string

func sourcesHandler(c *gin.Context) {
	var sourcePage Page
	sourcePage.Caption = packCaption("Sources")
	sourcePage.Body = createSourcePage()
	sourcePage.Tmpl = nil
	c.HTML(http.StatusOK, "sources.tmpl", gin.H{"Caption": packCaption("Sources")})
	//r.HTML(http.StatusOK, "index", sourcePage)
}

func createGrpSource() string {
	lblCaption = htmlLabel("caption", "caption:")
	edtCaption = htmlInput("caption", "caption")

	lblYear = htmlLabel("year", "year:")
	edtYear = htmlInput("year", "year:")

	lblShortName = htmlLabel("shortName", "shortName:")
	edtShortName = htmlInput("shortName", "shortName")
	controls := htmlAddNewLine(lblCaption+edtCaption) + "\n"
	controls += htmlAddNewLine(htmlIndent(lblYear+edtYear)) + "\n"
	controls += htmlAddNewLine(htmlIndent(lblShortName + edtShortName))
	grpSource = htmlGroupBox("source", controls)
	return grpSource
}
func createSourcePage() string {
	grpSource = createGrpSource()
	controls := grpSource
	return controls
}
