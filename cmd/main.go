package main

import (
	"log"
	"majipulse/api"
	"majipulse/blockchain"
	"majipulse/config"
	"net/http"
)

func main() {

	var bc = blockchain.NewBlockchain()

	// Load configuration
	cfg := config.LoadConfig()

	// Set up routes
	router := api.SetupRouter(bc)

	// Start the server
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
