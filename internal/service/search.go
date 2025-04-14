package service

import (
	"html/template"
	"net/http"
)

type TemplateData struct {
	Results []string
}

func render(w http.ResponseWriter, name string, data any) {
	tmpl := template.Must(template.ParseGlob("internal/templates/*.tmpl"))
	tmpl.ExecuteTemplate(w, name, data)
}

func (service *Service) homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home", nil)
}

func (service *Service) searchHandler(w http.ResponseWriter, r *http.Request) {
	isHtmx := r.Header.Get("HX-Request")
	searchResult := r.URL.Query().Get("search")
	if isHtmx == "true" {
		render(w, "results", []string{searchResult})
	} else {
		render(w, "home", TemplateData{Results: []string{searchResult}})
	}
}
