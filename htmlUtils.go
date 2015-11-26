package main

type HTML struct {
	lines []string
}

func (html HTML) String() string {
	strOut := ""
	for _, curLine := range html.lines {
		strOut += curLine
	}
	return strOut
}
func (html *HTML) addLine(line string) {
	html.lines = append(html.lines, line)
}

func packCaption(caption string) string {
	var html HTML
	html.addLine(" - " + caption)
	return html.String()
}

func htmlHeader(withOutHeader string) string {
	var html HTML
	html.addLine("<h2>" + withOutHeader + "</h2>")
	return html.String()
}

func htmlLink(path string, caption string) string {
	var html HTML
	html.addLine("<a href=\"" + path + "\">" + caption + "</a>")
	return html.String()
}

func htmlLabel(name string, caption string) string {
	var html HTML
	html.addLine("<label for=\"" + name + "\">" + caption + "</label>")
	return html.String()
}

func htmlGroupBox(caption string, controls string) string {
	var html HTML
	html.addLine("<fieldset>" + "\n")
	html.addLine(htmlIndent("<legend>"+caption+"</legend>") + "\n")
	html.addLine(htmlIndent(controls) + "\n")
	html.addLine("</fieldset>")
	return html.String()
}

func htmlTable(rows []string) string {
	var html HTML
	html.addLine("<table>\n")
	for _, curRow := range rows {
		html.addLine(htmlIndent(curRow) + "\n")
	}
	html.addLine("</table>")
	return html.String()
}

func htmlTableRow(row string) string {
	var html HTML
	html.addLine("<tr>" + "\n")
	html.addLine(htmlIndent(row) + "\n")
	html.addLine("</tr>")
	return html.String()
}

func htmlTableColumn(row string) string {
	var html HTML
	html.addLine("<td>" + row + "</td>")
	return html.String()
}

func htmlInput(name string, caption string) string {
	var html HTML
	html.addLine("<input type=\"text\">")
	return html.String()
}

func htmlIndent(indent string) string {
	var html HTML
	html.addLine("  " + indent)
	return html.String()
}

func htmlAddNewLine(curLine string) string {
	var html HTML
	html.addLine(curLine + "<br>")
	return html.String()
}
