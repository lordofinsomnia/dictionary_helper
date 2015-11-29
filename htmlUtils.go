package main

import (
	"fmt"
	"strconv"
)

type HTML struct {
	lines []string
}

func (html *HTML) String() string {
	strOut := ""
	for _, curLine := range html.lines {
		strOut += curLine
	}
	return strOut
}

func (html *HTML) addLineWithoutLF(line string) {
	html.lines = append(html.lines, line)
}

func (html *HTML) addLine(line string) {
	html.addLineWithoutLF(line + "\n")
}

func htmlHeader(withOutHeader string) string {
	var html HTML
	html.addLineWithoutLF("<h2>" + withOutHeader + "</h2>")
	return html.String()
}

func htmlLink(path string, caption string) string {
	var html HTML
	html.addLineWithoutLF("<a href=\"" + path + "\">" + caption + "</a>")
	return html.String()
}

func htmlLabel(name string, caption string) string {
	var html HTML
	html.addLineWithoutLF("<label for=\"" + name + "\">" + caption + "</label>")
	return html.String()
}

func htmlGroupBox(caption string, controls HTML) (string, HTML) {
	var html HTML
	html.addLine("<fieldset>")
	html.addLine(htmlIndent("<legend>" + caption + "</legend>"))
	for _, curRow := range controls.lines {
		html.addLineWithoutLF(htmlIndent(curRow))
	}
	html.addLine("</fieldset>")
	return html.String(), html
}

func htmlTable(htmls []HTML) (string, HTML) {
	var html HTML
	html.addLine("<table>")
	for _, curHtml := range htmls {
		for _, curRow := range curHtml.lines {
			html.addLineWithoutLF(htmlIndent(curRow))
		}
	}
	html.addLine("</table>")

	return html.String(), html
}

func htmlTableRow(row HTML) (string, HTML) {
	var html HTML
	html.addLine("<tr>")
	for _, curRow := range row.lines {
		html.addLine(curRow)
	}
	html.addLine("</tr>")
	return html.String(), html
}

func htmlTableColumn(row string) string {
	var html HTML
	html.addLineWithoutLF("<td>" + row + "</td>")
	return html.String()
}

func htmlInput(name string, caption string) string {
	var html HTML
	html.addLineWithoutLF("<input type=\"text\">")
	return html.String()
}

func htmlIndent(indent string) string {
	var html HTML
	html.addLineWithoutLF("  " + indent)
	return html.String()
}

func (html *HTML) htmlIndent3() {
	for i, _ := range html.lines {
		html.lines[i] = "  " + html.lines[i]
	}
}

func htmlIndent2(indent HTML) HTML {
	var html HTML
	fmt.Println("before indent:")
	fmt.Println(indent.String())

	for i, curLine := range indent.lines {
		fmt.Printf("i: %i, %s", i, curLine)
		html.addLine("--" + curLine)
	}
	fmt.Println("after indent:")
	fmt.Println(html.String())

	return html
}

func (html *HTML) htmlDump(caption string) {
	fmt.Println("----------------------")
	fmt.Println("html " + caption + " dump string:")
	fmt.Println(html.String())
	fmt.Println("----------------------")

	fmt.Println("html " + caption + " raw dump:")
	for i, curLine := range html.lines {
		fmt.Printf("%s:%s", strconv.Itoa(i), curLine)
	}
	fmt.Println("----------------------")
}

func htmlAddNewLine(curLine string) string {
	var html HTML
	html.addLineWithoutLF(curLine + "<br>")
	return html.String()
}
