package models

import (
	"majipulse/blockchain"
	"time"

	"github.com/google/uuid"
)

type WastageData struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
}

// ReportWastage logs the wastage data to the blockchain.
func ReportWastage(bc *blockchain.Blockchain, wastageData WastageData) error {
	transaction := blockchain.Transaction{
		ID:          uuid.New().String(),
		UserID:      wastageData.UserID,
		Description: wastageData.Description,
		Timestamp:   time.Now().String(),
	}
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)
	return nil
}
