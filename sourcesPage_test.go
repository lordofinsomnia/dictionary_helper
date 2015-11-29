package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSources(t *testing.T) {
	webApp = nil
	Convey("StartServer", t, func() {
		configureServer()
		response := getResponse(webApp, "GET", "/sources")
		responseStr := response.Body.String()
		Convey("Sources works", func() {
			Convey("Has all gui items", func() {
				createSourcePage()
				Convey("Has source groupbox", func() {
					Convey("Has groupbox caption", func() {
						So(responseStr, ShouldContainSubstring, grpSource)
					})
					Convey("Has caption", func() {
						Convey("Has caption label", func() {
							So(responseStr, ShouldContainSubstring, lblCaption)
						})
						Convey("Has caption editbox", func() {
							So(responseStr, ShouldContainSubstring, edtCaption)
						})
					})
					Convey("Has year", func() {
						Convey("Has year label", func() {
							So(responseStr, ShouldContainSubstring, lblYear)
						})
						Convey("Has year editbox", func() {
							So(responseStr, ShouldContainSubstring, edtYear)
						})
					})
					Convey("Has shortname", func() {
						Convey("Has shortname label", func() {
							So(responseStr, ShouldContainSubstring, lblShortName)
						})
						Convey("Has shortname editbox", func() {
							So(responseStr, ShouldContainSubstring, edtShortName)
						})
					})
				})
				Convey("Has buttons", func() {
					Convey("Has button add", func() {
						So(responseStr, ShouldContainSubstring, btnAddSource)
					})
					Convey("Has button edit", func() {
						So(responseStr, ShouldContainSubstring, btnEditSource)
					})
					Convey("Has button delete", func() {
						So(responseStr, ShouldContainSubstring, btnDeleteSource)
					})
					Convey("Has button export", func() {
						So(responseStr, ShouldContainSubstring, btnExportSource)
					})
					Convey("Has button import", func() {
						So(responseStr, ShouldContainSubstring, btnImportSource)
					})
					Convey("Has button print", func() {
						So(responseStr, ShouldContainSubstring, btnPrintSource)
					})
				})
			})
		})
	})
}
