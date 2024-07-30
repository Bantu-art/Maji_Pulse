package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SensorData struct to hold the sensor data
type SensorData struct {
	Area           string  `json:"area"`
	FlowRate       float64 `json:"flow_rate"`
	Leakage        bool    `json:"leakage"`
	FairPercentage float64 `json:"fair_percentage"`
}

// Handle WebSocket connections
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	leakageTicker := time.NewTicker(2 * time.Minute)
	defer leakageTicker.Stop()

	for {
		select {
		case <-time.After(1 * time.Second):
			// Simulate flow rate and fair distribution data every second
			data := []SensorData{
				{"Area 1", rand.Float64() * 100, false, rand.Float64() * 100},
				{"Area 2", rand.Float64() * 100, false, rand.Float64() * 100},
				{"Area 3", rand.Float64() * 100, false, rand.Float64() * 100},
			}
			if err := conn.WriteJSON(data); err != nil {
				fmt.Println("Error writing JSON to WebSocket:", err)
				return
			}

		case <-leakageTicker.C:
			// Simulate leakage data every 2 minutes
			data := []SensorData{
				{"Area 1", 0, rand.Intn(2) == 1, 0},
				{"Area 2", 0, rand.Intn(2) == 1, 0},
				{"Area 3", 0, rand.Intn(2) == 1, 0},
			}
			if err := conn.WriteJSON(data); err != nil {
				fmt.Println("Error writing JSON to WebSocket:", err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))

	fmt.Println("Server started at :8090")
	http.ListenAndServe(":8090", nil)
}
