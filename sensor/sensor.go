package sensor

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type SensorData struct {
	SensorID  string  `json:"sensor_id"`
	FlowRate  float64 `json:"flow_rate"` // Simulated flow rate in liters per minute
	Timestamp string  `json:"timestamp"`
}

func generateSensorData() SensorData {
	return SensorData{
		SensorID:  "sensor_1",
		FlowRate:  rand.Float64() * 10, // Random flow rate between 0 and 10 L/min
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func sendSensorData(data SensorData) {
	url := "http://localhost:8080/api/sensor-data" // Replace with your actual endpoint
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ticker := time.NewTicker(5 * time.Second) // Send data every 5 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sensorData := generateSensorData()
			sendSensorData(sensorData)
			println("Sent sensor data:", sensorData)
		}
	}
}
