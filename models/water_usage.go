package models

import (
	"majipulse/blockchain"
	"time"
)

// WaterUsage represents a water usage record.
type WaterUsage struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// GetWaterUsage retrieves water usage records for a specific user from the blockchain.
func GetWaterUsage(bc *blockchain.Blockchain, userID string) ([]WaterUsage, error) {
	var waterUsages []WaterUsage

	// Iterate through the blockchain and fetch water usage records for the specified user
	for _, block := range bc.GetBlocks() {
		for _, transaction := range block.Transactions {
			if transaction.UserID == userID {
				waterUsages = append(waterUsages, WaterUsage{
					UserID:    transaction.UserID,
					Amount:    transaction.Amount,
					Timestamp: time.Now(), // You may want to parse the timestamp from the transaction
				})
			}
		}
	}

	return waterUsages, nil
}
