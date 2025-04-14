package service

import (
	"log"
	"net/http"
)

type Service struct {
}

func NewService() Service {
	service := Service{}

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("GET /", service.homeHandler)
	mux.HandleFunc("GET /search", service.searchHandler)

	addr := ":8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

	return service
}
