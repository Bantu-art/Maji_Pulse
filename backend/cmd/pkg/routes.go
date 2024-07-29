package handler

import (
	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the API routes
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Health check route
	router.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")

	// Water usage routes
	router.HandleFunc("/api/water-usage", GetWaterUsageHandler).Methods("GET")
	router.HandleFunc("/api/water-usage", UpdateWaterUsageHandler).Methods("POST")

	// Leak detection route
	router.HandleFunc("/api/leak-detection", DetectLeakHandler).Methods("POST")

	// User settings routes
	router.HandleFunc("/api/users/{id}", GetUserSettingsHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", UpdateUserSettingsHandler).Methods("PUT")

	return router
}
