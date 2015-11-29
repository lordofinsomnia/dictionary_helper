package main

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
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

func packCaption(caption string) string {
	return " - " + caption
}

func sourcesHandler(c *gin.Context) {
	var sourcePage Page
	sourcePage.Caption = packCaption("Sources")
	sourcePage.Body = createSourcePage()
	sourcePage.Tmpl = "templates/sources.tmpl"
	c.HTML(http.StatusOK, "sources", gin.H{"Caption": packCaption("Sources")})
}

func createGrpSource() string {
	lblCaption = htmlTableColumn(htmlLabel("caption", "caption:"))
	edtCaption = htmlTableColumn(htmlInput("caption", "caption"))
	var rowCaption HTML
	rowCaption.addLineWithoutLF(lblCaption)
	rowCaption.addLineWithoutLF(edtCaption)
	rowCaption.htmlIndent3()
	rowCaption.htmlDump("rowCaption")

	lblYear = htmlTableColumn(htmlLabel("year", "year:"))
	edtYear = htmlTableColumn(htmlInput("year", "year:"))

	var rowYear HTML
	rowYear.addLineWithoutLF(lblYear)
	rowYear.addLineWithoutLF(edtYear)
	rowYear.htmlIndent3()
	rowYear.htmlDump("rowYear")

	lblShortName = htmlTableColumn(htmlLabel("shortName", "shortName:"))
	edtShortName = htmlTableColumn(htmlInput("shortName", "shortName"))

	var rowShortName HTML
	rowShortName.addLineWithoutLF(lblShortName)
	rowShortName.addLineWithoutLF(edtShortName)
	rowShortName.htmlIndent3()
	rowShortName.htmlDump("rowShortName")

	var htmls []HTML

	_, captionHtml := htmlTableRow(rowCaption)
	_, rowHtml := htmlTableRow(rowYear)
	_, shortNameHtml := htmlTableRow(rowShortName)
	captionHtml.htmlDump("captionHtml")
	rowHtml.htmlDump("rowHtml")
	shortNameHtml.htmlDump("shortNameHtml")

	htmls = make([]HTML, 3)
	htmls[0] = captionHtml
	htmls[1] = rowHtml
	htmls[2] = shortNameHtml

	_, tableHtml := htmlTable(htmls)
	_, grpSourceHtml := htmlGroupBox("source", tableHtml)

	return grpSourceHtml.String()
}
func createSourcePage() string {
	grpSource = createGrpSource()
	controls := grpSource
	return controls
}
