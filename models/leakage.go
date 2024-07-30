package models

import (
	"majipulse/blockchain"
	"time"

	"github.com/google/uuid"
)

// LeakageData represents data for a leakage detection.
type LeakageData struct {
	SensorID string  `json:"sensor_id"`
	FlowRate float64 `json:"flow_rate"`
}

// ReportLeakage logs the leakage data to the blockchain.
func ReportLeakage(bc *blockchain.Blockchain, leakageData LeakageData) error {
	transaction := blockchain.Transaction{
		ID:        uuid.New().String(),
		SensorID:  leakageData.SensorID,
		FlowRate:  leakageData.FlowRate,
		Timestamp: time.Now().String(),
	}
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)
	return nil
}
