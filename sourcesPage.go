package main

import (
	"github.com/martini-contrib/render"
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

func sourcesHandler(r render.Render) {
	var sourcePage Page
	sourcePage.Caption = packCaption("Sources")
	sourcePage.Body = createSourcePage()
	sourcePage.Tmpl = nil

	r.HTML(http.StatusOK, "index", sourcePage)
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
	return grpSource
}
