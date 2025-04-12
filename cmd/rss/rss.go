package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, data any, pages ...string) {
	files := append([]string{"internal/templates/base.tmpl"}, pages...)
	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.ExecuteTemplate(w, "base", data)
}

func homeGET(w http.ResponseWriter, r *http.Request) {
	render(w, nil, "internal/templates/home.tmpl")
}

func homePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Posted to home page")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("GET /", homeGET)
	mux.HandleFunc("POST /", homePOST)

	addr := ":8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
