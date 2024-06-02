package main

import (
	"log"
	"net/http"

	"github.com/swadand/ebook-creator-backend/pkg/api"
	"github.com/swadand/ebook-creator-backend/pkg/db"
)

func main() {
	r := api.SetupRoutes()
	db.Init()

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}

	log.Printf("Starting server at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
