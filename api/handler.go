package api

import (
	"majipulse/blockchain" // Import the blockchain package
	"majipulse/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Global blockchain instance
var bc = blockchain.NewBlockchain()

// CORSMiddleware handles CORS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func HomePageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// LeakageDetectorHandler handles leakage detection requests
func LeakageDetectorHandler(c *gin.Context) {
	var leakageData models.LeakageData
	if err := c.ShouldBindJSON(&leakageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid leakage data"})
		return
	}

	// Log leakage to the blockchain
	transaction := blockchain.Transaction{
		ID:        uuid.New().String(),
		SensorID:  leakageData.SensorID,
		FlowRate:  leakageData.FlowRate,
		Timestamp: time.Now().String(),
	}

	// Add the transaction to the blockchain
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)

	c.JSON(http.StatusOK, gin.H{"message": "Leakage detected and reported"})
}

// ViewWaterUsageHandler allows users to view their water usage
func ViewWaterUsageHandler(c *gin.Context) {
	userID := c.Param("user_id")
	waterUsage, err := models.GetWaterUsage(bc, userID) // Adjusted to use blockchain
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, waterUsage)
}

// ReportWastageHandler allows users to report water wastage
func ReportWastageHandler(c *gin.Context) {
	var wastageData models.WastageData
	if err := c.ShouldBindJSON(&wastageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wastage data"})
		return
	}

	// Create a transaction for the wastage event
	transaction := blockchain.Transaction{
		ID:          uuid.New().String(),
		UserID:      wastageData.UserID,
		Description: wastageData.Description,
		Timestamp:   time.Now().String(),
	}

	// Add the transaction to the blockchain
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)

	c.JSON(http.StatusOK, gin.H{"message": "Wastage reported successfully"})
}

// HandleSensorData handles incoming sensor data
func HandleSensorData(c *gin.Context) {
	var sensorData models.SensorData
	if err := c.ShouldBindJSON(&sensorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor data"})
		return
	}

	// Log sensor data to the blockchain
	transaction := blockchain.Transaction{
		ID:        uuid.New().String(),
		SensorID:  sensorData.SensorID,
		FlowRate:  sensorData.FlowRate,
		Timestamp: time.Now().String(),
	}

	// Add the transaction to the blockchain
	bc.AddBlock([]blockchain.Transaction{transaction}, bc.GetBlocks()[len(bc.GetBlocks())-1].Hash)

	c.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}
