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

func sourcesHandler(c *gin.Context) {
	var sourcePage Page
	sourcePage.Caption = packCaption("Sources")
	sourcePage.Body = createSourcePage()
	sourcePage.Tmpl = "templates/sources.tmpl"
	c.HTML(http.StatusOK, "sources", gin.H{"Caption": packCaption("Sources")})
}

func createGrpSource() string {
	/*lblCaption = htmlLabel("caption", "caption:")
	edtCaption = htmlInput("caption", "caption")

	lblYear = htmlLabel("year", "year:")
	edtYear = htmlInput("year", "year:")

	lblShortName = htmlLabel("shortName", "shortName:")
	edtShortName = htmlInput("shortName", "shortName")
	controls := htmlAddNewLine(lblCaption+edtCaption) + "\n"
	controls += htmlAddNewLine(htmlIndent(lblYear+edtYear)) + "\n"
	controls += htmlAddNewLine(htmlIndent(lblShortName + edtShortName))
	grpSource = htmlGroupBox("source", controls)*/

	lblCaption = htmlTableColumn(htmlLabel("caption", "caption:"))
	edtCaption = htmlTableColumn(htmlInput("caption", "caption"))
	var rowCaption HTML
	rowCaption.addLineWithoutLF(lblCaption)
	rowCaption.addLineWithoutLF(edtCaption)
	rowCaption.htmlIndent3()
	rowCaption.htmlDump("rowCaption")

	//columnCaption := htmlTableColumn(rowCaption.String()) + "e1te"

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

	/*captionHtml.htmlIndent3()
	rowHtml.htmlIndent3()
	shortNameHtml.htmlIndent3()*/

	htmls = make([]HTML, 3)
	htmls[0] = captionHtml
	htmls[1] = rowHtml
	htmls[2] = shortNameHtml

	captionHtml.htmlDump("captionHtml")
	_, tableHtml := htmlTable(htmls)
	tableHtml.htmlDump("tableHtml")

	_, grpSourceHtml := htmlGroupBox("source", tableHtml)
	grpSourceHtml.htmlDump("grpSourceHtml")

	/*grpHtml.addLine()
	grpHtml.addLine(htmlTableRow(rowYear.String()))
	grpHtml.addLine(htmlTableRow(rowShortName.String()))
	grpHtml.htmlDump("grpHtml")*/
	//grpSourceStr, grpSourceHtml := htmlGroupBox("source", controls)

	/*var controls HTML
	controls.addLine(htmlTableColumn(lblCaption))
	controls.addLine(htmlTableColumn(edtCaption))
	controls.addLine(htmlAddNewLine(htmlIndent(lblYear + edtYear)))
	controls.addLineWithoutLF(htmlAddNewLine(htmlIndent(lblShortName + edtShortName)))*/
	/*htmlIndent(columnCaption), htmlIndent(rowYear.String())*/
	//var rows []string
	//rows := make([]string, 2)
	/*controlsStr, controlsHTML := "htmlTable(rows)"
	grpSourceStr, grpSourceHtml := htmlGroupBox("source", controls)
	fmt.Println("controlsStr:")
	fmt.Println(controlsStr)

	fmt.Println("controlsHTML:")
	fmt.Println(controlsHTML)

	fmt.Println("grpSourceStr:")
	fmt.Println(grpSourceStr)

	fmt.Println("grpSourceHtml:")
	fmt.Println(grpSourceHtml)

	fmt.Println("grpSource:")
	fmt.Println(htmlIndent2(grpSourceHtml))*/

	return grpSourceHtml.String()
}
func createSourcePage() string {
	grpSource = createGrpSource()
	controls := grpSource
	return controls
}
