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
	return "<fieldset><legend>" + caption + "</legend>" + controls + "</fieldset>"
}

func htmlInput(name string, caption string) string {
	return "<input type=\"text\">" + caption + "</input>"
}

func htmlAddNewLine(html string) string {
	return html + "<br>"
}
