package main

var lblCaption string
var edtCaption string
var lblYear string
var edtYear string
var lblShortName string
var edtShortName string
var controls string
var grpSource string

func createGrpSource() string {
	lblCaption = htmlLabel("caption", "caption")
	edtCaption = htmlAddNewLine(htmlInput("caption", "caption"))
	lblYear = htmlLabel("year", "year")
	edtYear = htmlAddNewLine(htmlInput("year", "year"))
	lblShortName = htmlLabel("shortName", "shortName")
	edtShortName = htmlAddNewLine(htmlInput("shortName", "shortName"))
	controls = lblCaption + edtCaption + lblYear + edtYear + lblShortName + edtShortName
	grpSource = htmlGroupBox("source", controls)
	return grpSource
}
func createSourcePage() string {
	grpSource = createGrpSource()
	return grpSource
}
