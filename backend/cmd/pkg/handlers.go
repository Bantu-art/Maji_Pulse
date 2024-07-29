package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetWaterUsageHandler retrieves the current water usage data
func GetWaterUsageHandler(w http.ResponseWriter, r *http.Request) {
	// Mock data for demonstration
	usageData := map[string]interface{}{
		"currentUsage": 150, // in liters
		"averageUsage": 120, // in liters
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usageData)
}

// UpdateWaterUsageHandler updates water usage data
func UpdateWaterUsageHandler(w http.ResponseWriter, r *http.Request) {
	var usageData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&usageData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would typically update the usage data in the database

	w.WriteHeader(http.StatusNoContent) // No content to return
}

// DetectLeakHandler handles leak detection requests
func DetectLeakHandler(w http.ResponseWriter, r *http.Request) {
	// Mock leak detection logic
	leakDetected := true // Simulate leak detection

	if leakDetected {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Leak detected!"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "No leaks detected."})
	}
}

// GetUserSettingsHandler retrieves user settings
func GetUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	vars := mux.Vars(r)
	userID := vars["id"]

	// Mock user settings data
	userSettings := map[string]interface{}{
		"userID":    userID,
		"threshold": 100, // in liters
		"alerts":    true,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userSettings)
}

// UpdateUserSettingsHandler updates user settings
func UpdateUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	// vars := mux.Vars(r)
	// userID := vars["id"]

	var settings map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent) // No content to return
}
