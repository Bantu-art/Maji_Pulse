package models

import (
	"majipulse/blockchain"
	"time"

	"github.com/google/uuid"
)

type SensorData struct {
	SensorID string  `json:"sensor_id"`
	FlowRate float64 `json:"flow_rate"`
}

// SaveSensorData logs sensor data to the blockchain.
func SaveSensorData(bc *blockchain.Blockchain, sensorData SensorData) error {
	transaction := blockchain.Transaction{
		ID:        uuid.New().String(),
		SensorID:  sensorData.SensorID,
		FlowRate:  sensorData.FlowRate,
		Timestamp: time.Now().String(),
	}
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)
	return nil
}
