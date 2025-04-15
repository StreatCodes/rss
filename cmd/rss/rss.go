package main

import (
	"log"

	"github.com/streatCodes/rss/internal/service"
)

func main() {
	_, err := service.New("rss.db")
	if err != nil {
		log.Fatalf("Error initiating service %s\n", err)
	}
}
