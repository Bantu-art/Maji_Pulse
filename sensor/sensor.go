package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// SensorData represents the data structure for the sensor readings.
type SensorData struct {
	SensorID string  `json:"sensor_id"`
	FlowRate float64 `json:"flow_rate"`
}

func main() {
	// Set the sensor ID
	sensorID := "sensor_001"

	// Run the sensor simulator in an infinite loop
	for {
		// Generate a random flow rate between 0 and 2.0
		flowRate := rand.Float64() * 2.0

		// Create a new SensorData instance
		sensorData := SensorData{
			SensorID: sensorID,
			FlowRate: flowRate,
		}

		// Send the sensor data to the API
		err := sendSensorData(sensorData)
		if err != nil {
			log.Println("Error sending sensor data:", err)
		}

		// Wait for a random interval between 1 and 5 seconds
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}

// sendSensorData sends the sensor data to the specified API endpoint.
func sendSensorData(data SensorData) error {
	// Convert the sensor data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Define the API endpoint (replace with your actual endpoint)
	apiEndpoint := "http://localhost:8080/api/sensor"

	// Send a POST request to the API
	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send data, status code: %d", resp.StatusCode)
	}

	log.Printf("Sensor data sent: %+v", data)
	return nil
}
