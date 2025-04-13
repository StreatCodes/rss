package main

import (
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Results []string
}

func render(w http.ResponseWriter, name string, data any) {
	tmpl := template.Must(template.ParseGlob("internal/templates/*.tmpl"))
	tmpl.ExecuteTemplate(w, name, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	isHtmx := r.Header.Get("HX-Request")
	searchResult := r.URL.Query().Get("search")
	if isHtmx == "true" {
		render(w, "results", []string{searchResult})
	} else {
		render(w, "home", TemplateData{Results: []string{searchResult}})
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /search", searchHandler)

	addr := ":8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
