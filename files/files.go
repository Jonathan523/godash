package files

import (
	"html/template"
	"net/http"
)

const TemplatesFolder = "templates/"

func ParseAndServeHtml(w http.ResponseWriter, htmlFile string, info interface{}) {
	parsedHtml, _ := template.ParseFiles(TemplatesFolder+htmlFile, TemplatesFolder+"_base.gohtml")
	_ = parsedHtml.Execute(w, info)
}
