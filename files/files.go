package files

import (
	"html/template"
	"net/http"
)

const templatesFolder = "templates/"

func ParseHtml(w http.ResponseWriter, htmlFile string, info interface{}) {
	parsedHtml, _ := template.ParseFiles(templatesFolder+htmlFile, templatesFolder+"_base.gohtml")
	_ = parsedHtml.Execute(w, info)
}
