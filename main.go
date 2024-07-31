package main

import (
	"log"
	"majipulse/api"
	"majipulse/blockchain"
	"majipulse/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up blockchain instance
	bc := blockchain.NewBlockchain()

	// Set up routes
	router := api.SetupRouter(bc)

	// Serve frontend files from the "frontend" directory
	router.Static("/frontend", "./frontend") // Adjust path as necessary
	router.LoadHTMLGlob("frontend/*.html")   // Load all HTML files in the frontend directory

	// Start the server
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
