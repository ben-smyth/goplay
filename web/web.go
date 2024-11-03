package web

import (
	"html/template"
	"path/filepath"
)

func HandleStatic() {

}

func PopulateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "web/templates"
	template.Must(result.ParseGlob(filepath.Join(basePath, "components/*.gohtml")))
	template.Must(result.ParseGlob(filepath.Join(basePath, "pages/*.gohtml")))
	return result
}
