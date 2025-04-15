package service

import (
	"log"
	"net/http"
	"time"

	"github.com/streatCodes/rss/internal/db"
	bolt "go.etcd.io/bbolt"
)

type Service struct {
	db *db.DB
}

func New(dbPath string) (*Service, error) {
	db, err := db.New(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	service := &Service{db: db}

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	mux.HandleFunc("GET /", service.homeHandler)
	mux.HandleFunc("GET /search", service.searchHandler)

	addr := ":8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

	return service, nil
}
