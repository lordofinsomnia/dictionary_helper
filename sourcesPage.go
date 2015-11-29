package main

import (
	//	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	lblCaption      string
	edtCaption      string
	lblYear         string
	edtYear         string
	lblShortName    string
	edtShortName    string
	controls        string
	buttons         string
	grpSource       string
	btnAddSource    string
	btnEditSource   string
	btnDeleteSource string
	btnExportSource string
	btnImportSource string
	btnPrintSource  string
)

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

	lblYear = htmlTableColumn(htmlLabel("year", "year:"))
	edtYear = htmlTableColumn(htmlInput("year", "year:"))

	var rowYear HTML
	rowYear.addLineWithoutLF(lblYear)
	rowYear.addLineWithoutLF(edtYear)
	rowYear.htmlIndent3()

	lblShortName = htmlTableColumn(htmlLabel("shortName", "shortName:"))
	edtShortName = htmlTableColumn(htmlInput("shortName", "shortName"))

	var rowShortName HTML
	rowShortName.addLineWithoutLF(lblShortName)
	rowShortName.addLineWithoutLF(edtShortName)
	rowShortName.htmlIndent3()

	var htmls []HTML

	_, captionHtml := htmlTableRow(rowCaption)
	_, rowHtml := htmlTableRow(rowYear)
	_, shortNameHtml := htmlTableRow(rowShortName)

	htmls = make([]HTML, 3)
	htmls[0] = captionHtml
	htmls[1] = rowHtml
	htmls[2] = shortNameHtml

	_, tableHtml := htmlTable(htmls)
	_, grpSourceHtml := htmlGroupBox("source", tableHtml)

	return grpSourceHtml.String()
}

func createButtons() string {
	btnAddSource = htmlButton("add")
	btnEditSource = htmlButton("edit")
	btnDeleteSource = htmlButton("delete")
	btnExportSource = htmlButton("export")
	btnImportSource = htmlButton("import")
	btnPrintSource = htmlButton("print")
	return ""
}

func createSourcePage() string {
	grpSource = createGrpSource()
	buttons = createButtons()

	controls := grpSource
	controls += buttons
	return controls
}
