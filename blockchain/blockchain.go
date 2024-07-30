package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"
)

// Transaction represents a record of an event in the blockchain.
type Transaction struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id,omitempty"`     // For water usage and wastage
	SensorID    string  `json:"sensor_id,omitempty"`   // For leakages and sensor data
	FlowRate    float64 `json:"flow_rate,omitempty"`   // For leakages and sensor data
	Amount      float64 `json:"amount,omitempty"`      // For water usage
	Description string  `json:"description,omitempty"` // For wastage
	Timestamp   string  `json:"timestamp"`
}

// Block represents a single block in the blockchain.
type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Hash         string        `json:"hash"`
}

// Blockchain represents the entire blockchain.
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates a new blockchain with the genesis block.
func NewBlockchain() *Blockchain {
	b := &Blockchain{}
	b.AddBlock([]Transaction{}, "") // Create the genesis block
	return b
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transactions []Transaction, prevHash string) {
	block := &Block{
		Index:        len(bc.blocks) + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Hash:         "",
	}
	block.Hash = block.calculateHash()
	bc.blocks = append(bc.blocks, block)
}

// calculateHash calculates the hash of the block.
func (b *Block) calculateHash() string {
	transactionsJSON, _ := json.Marshal(b.Transactions)
	record := strconv.Itoa(b.Index) + b.Timestamp + b.PrevHash + string(transactionsJSON)
	h := sha256.New()
	h.Write([]byte(record))
	return string(h.Sum(nil))
}

// GetBlocks returns all blocks in the blockchain.
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}
