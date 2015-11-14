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
	html += "<legend>" + caption + "</legend>\n"
	html += controls
	html += "</fieldset>"
	return html
}

func htmlInput(name string, caption string) string {
	return "<input type=\"text\">"
}

func htmlAddNewLine(html string) string {
	return html + "<br>"
}
