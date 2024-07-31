package api

import (
	"majipulse/blockchain"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router and sets up the routes.
func SetupRouter(bc *blockchain.Blockchain) *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Define your routes here
	router.GET("/", HomePageHandler)
	router.POST("/api/leakage", LeakageDetectorHandler)
	router.POST("/api/wastage", ReportWastageHandler)
	router.POST("/api/sensor", HandleSensorData)
	router.GET("/api/water-usage/:user_id", ViewWaterUsageHandler)

	return router
}
