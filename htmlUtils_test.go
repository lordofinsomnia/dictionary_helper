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
		Convey("htmlTableColumn func", func() {
			name := "test"
			expectedTableColumn := "<td>test</td>"
			So(htmlTableColumn(name), ShouldEqual, expectedTableColumn)
		})
		Convey("htmlTableRow func", func() {
			var name HTML
			name.addLineWithoutLF("test")
			name.htmlIndent3()
			expectedTableRow := "<tr>" + "\n"
			expectedTableRow += name.String() + "\n"
			expectedTableRow += "</tr>\n"
			gotTableRowStr, gotTableRowHtml := htmlTableRow(name)
			gotTableRowHtml.htmlDump("gotTableRowHtml")
			So(gotTableRowStr, ShouldEqual, expectedTableRow)
		})
		Convey("htmlTable func", func() {

			var tableRows []string
			tableRows = make([]string, 2)
			var htmlRows HTML

			htmlRows.addLine("<tr>")
			htmlRows.addLine(htmlIndent("<td>t00</td><td>t01</td><td>t02</td>"))
			htmlRows.addLine("</tr>")
			htmlRows.addLine("<tr>")
			htmlRows.addLine(htmlIndent("<td>t10</td><td>t11</td><td>t12</td>"))
			htmlRows.addLine("</tr>")

			tableRows[0] = "<tr>" + "\n"
			tableRows[0] += "  <td>t00</td><td>t01</td><td>t02</td>" + "\n"
			tableRows[0] += "</tr>"
			tableRows[1] = "<tr>" + "\n"
			tableRows[1] += "  <td>t10</td><td>t11</td><td>t12</td>" + "\n"
			tableRows[1] += "</tr>"

			expectedTable := "<table>" + "\n"
			expectedTable += "  <tr>" + "\n"
			expectedTable += "    <td>t00</td><td>t01</td><td>t02</td>" + "\n"
			expectedTable += "  </tr>" + "\n"
			expectedTable += "  <tr>" + "\n"
			expectedTable += "    <td>t10</td><td>t11</td><td>t12</td>" + "\n"
			expectedTable += "  </tr>" + "\n"
			expectedTable += "</table>\n"

			expectedTable2 := "<table>" + "\n"
			expectedTable2 += "</table>"

			htmlRows.htmlDump("htmlRows")
			var htmls []HTML
			htmls = make([]HTML, 1)
			htmls[0] = htmlRows

			_, gotTableHtml := htmlTable(htmls)

			So(gotTableHtml.String(), ShouldEqual, expectedTable)
		})
		Convey("htmlGroupBox func", func() {
			name := "test"
			caption := "test"
			expected := "<fieldset>\n"
			expected += htmlIndent("<legend>test</legend>\n")
			expected += htmlIndent("<label for=\"test\">test</label><br>" + "<input type=\"text\">" + "\n")
			expected += "</fieldset>\n"
			label := htmlLabel(name, caption)
			input := htmlInput(name, caption)
			//controls := label + "<br>" + input
			var controls HTML
			controls.addLine(label + "<br>" + input)
			htmlStr, _ := htmlGroupBox(name, controls)

			So(htmlStr, ShouldEqual, expected)
		})
	})
}
