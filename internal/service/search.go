package service

import (
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/streatCodes/rss/rss"
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

	log.Printf("Search: %s\n", searchResult)

	//If the search query is a URL then ingest the feed
	if parsedURL, err := url.ParseRequestURI(searchResult); err == nil {
		log.Println("Found URL")
		res, err := http.Get(searchResult)
		if err != nil {
			panic("TODO")
		}

		f, err := rss.Decode(res.Body)
		if err != nil {
			panic("TODO")
		}

		key := []byte(parsedURL.String())
		service.db.SaveFeed(key, &f.Channel)
	}

	if isHtmx == "true" {
		render(w, "results", []string{searchResult})
	} else {
		render(w, "home", TemplateData{Results: []string{searchResult}})
	}
}
