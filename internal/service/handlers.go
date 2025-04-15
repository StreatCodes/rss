package service

import (
	"html/template"
	"net/http"

	"github.com/streatCodes/rss/rss"
)

type TemplateData struct {
	Results []rss.Channel
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
	searchQuery := r.URL.Query().Get("search")

	channels, err := service.findChannel(searchQuery)
	if err != nil {
		panic("TODO")
	}

	if isHtmx == "true" {
		render(w, "results", channels)
	} else {
		render(w, "home", TemplateData{Results: channels})
	}
}
