package main

import (
	"config/config"
	"handler/api"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up routes
	router := api.SetupRouter()

	// Start the server
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
