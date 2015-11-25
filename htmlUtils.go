package main

func packCaption(caption string) string {
	return " - " + caption
}

func htmlHeader(withOutHeader string) string {
	return "<h2>" + withOutHeader + "</h2>"
}

func htmlLink(path string, caption string) string {
	return "<a href=\"" + path + "\">" + caption + "</a>"
}

func htmlLabel(name string, caption string) string {
	return "<label for=\"" + name + "\">" + caption + "</label>"
}

func htmlGroupBox(caption string, controls string) string {
	html := "<fieldset>\n"
	html += htmlIndent("<legend>" + caption + "</legend>" + "\n")
	html += htmlIndent(controls + "\n")
	html += "</fieldset>"
	return html
}

func htmlTable(rows []string) string {
	html := "<table>\n"
	for _, curRow := range rows {
		html += htmlIndent(curRow) + "\n"
	}
	html += "</table>"
	return html
}

func htmlTableRow(row string) string {
	html := "<tr>" + "\n"
	html += htmlIndent(row) + "\n"
	html += "</tr>"
	return html
}

func htmlTableColumn(row string) string {
	return "<td>" + row + "</td>"
}

func htmlInput(name string, caption string) string {
	return "<input type=\"text\">"
}

func htmlIndent(html string) string {
	return "  " + html
}

func htmlAddNewLine(html string) string {
	return html + "<br>"
}
