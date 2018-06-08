package controllers

import (
	"html/template"
	"net/http"
)

// IndexController - struct for IndexController
type IndexController struct {
}

var templates map[string]*template.Template

// compile new templates
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
}

func renderTemplate(w http.ResponseWriter, name string, template string) {
	// ensure the template exists in the map
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template doesn't exists!", http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Index show index page of web
func (i *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base")
}
