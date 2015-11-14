package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHTMLUtils(t *testing.T) {
	Convey("Testing utils func", t, func() {
		Convey("htmlHeader func", func() {
			test := "test"
			expected := "<h2>test</h2>"
			So(htmlHeader(test), ShouldEqual, expected)
		})
		Convey("htmlLink func", func() {
			path := "/test"
			caption := "test"
			expected := "<a href=\"/test\">test</a>"
			So(htmlLink(path, caption), ShouldEqual, expected)
		})
		Convey("htmlLabel func", func() {
			name := "test"
			caption := "test"
			expected := "<label for=\"test\">test</label>"
			So(htmlLabel(name, caption), ShouldEqual, expected)
		})
		Convey("htmlInput func", func() {
			name := "test"
			caption := "test"
			expected := "<input type=\"text\">"
			So(htmlInput(name, caption), ShouldEqual, expected)
		})
		Convey("packCaption func", func() {
			caption := "test"
			expected := " - test"
			So(packCaption(caption), ShouldEqual, expected)
		})
		Convey("htmlAddNewLine func", func() {
			name := "test"
			caption := "test"
			expectedLbl := "<input type=\"text\"><br>"
			So(htmlAddNewLine(""), ShouldEqual, "<br>")
			So(htmlAddNewLine(htmlInput(name, caption)), ShouldEqual, expectedLbl)
		})
		Convey("htmlGroupBox func", func() {
			name := "test"
			caption := "test"
			expected := "<fieldset>\n"
			expected += "<legend>test</legend>\n"
			expected += "<label for=\"test\">test</label><br>" + "<input type=\"text\">"
			expected += "</fieldset>"
			label := htmlLabel(name, caption)
			input := htmlInput(name, caption)
			controls := label + "<br>" + input
			So(htmlGroupBox(name, controls), ShouldEqual, expected)
		})
	})
}
