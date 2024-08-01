// backend/handlers.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// WaterUsage represents a water usage record
type WaterUsage struct {
	UserID    string `json:"user_id"`
	Amount    uint   `json:"amount"`
	Timestamp int64  `json:"timestamp"`
}

var usageData []WaterUsage

// recordUsageHandler handles the recording of water usage
func recordUsageHandler(w http.ResponseWriter, r *http.Request) {
	var usage WaterUsage
	if err := json.NewDecoder(r.Body).Decode(&usage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usage.Timestamp = time.Now().Unix()
	usageData = append(usageData, usage)
	log.Printf("Recorded usage: %+v\n", usage)
	w.WriteHeader(http.StatusCreated)
}

// getUsageHandler returns the recorded water usage data
func getUsageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(usageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
