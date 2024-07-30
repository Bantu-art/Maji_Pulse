package api

import (
	"Maji_pulse/models"
	"Maji_pulse/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LeakageDetectorHandler handles leakage detection requests
func LeakageDetectorHandler(c *gin.Context) {
	var leakageData models.LeakageData
	if err := c.ShouldBindJSON(&leakageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid leakage data"})
		return
	}

	err := services.DetectLeakage(leakageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leakage detected and reported"})
}

// ViewWaterUsageHandler allows users to view their water usage
func ViewWaterUsageHandler(c *gin.Context) {
	userID := c.Param("user_id")
	waterUsage, err := services.GetWaterUsageByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, waterUsage)
}

// ManageEqualDistributionHandler manages equal water distribution
func ManageEqualDistributionHandler(c *gin.Context) {
	var distributionData models.DistributionData
	if err := c.ShouldBindJSON(&distributionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid distribution data"})
		return
	}

	err := services.DistributeWaterEqually(distributionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Water distributed equally"})
}

// ReportWastageHandler allows users to report water wastage
func ReportWastageHandler(c *gin.Context) {
	var wastageData models.WastageData
	if err := c.ShouldBindJSON(&wastageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wastage data"})
		return
	}

	err := services.ReportWastage(wastageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wastage reported successfully"})
}

// HandleSensorData handles incoming sensor data
func HandleSensorData(c *gin.Context) {
	var sensorData models.SensorData
	if err := c.ShouldBindJSON(&sensorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor data"})
		return
	}

	// Process sensor data (e.g., detect leakage)
	if sensorData.FlowRate > 0 && sensorData.FlowRate < 0.1 {
		// Potential leakage detected
		reportLeakage(sensorData)
	}

	// Save sensor data to the database
	err := database.SaveSensorData(sensorData)
	if err != nil {
		utils.Error("Failed to save sensor data: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save sensor data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sensor data received"})
}
