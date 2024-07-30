package api

import (
    "majipulse/blockchain"
    "github.com/gin-gonic/gin"
    "majipulse/api"
)

// SetupRouter initializes the router and sets up the routes.
func SetupRouter(bc *blockchain.Blockchain) *gin.Engine {
    router := gin.Default()

    // Define your routes here
    router.POST("/api/leakage", api.LeakageDetectorHandler)
    router.POST("/api/wastage", api.ReportWastageHandler)
    router.POST("/api/sensor", api.HandleSensorData)
    router.GET("/api/water-usage/:user_id", api.ViewWaterUsageHandler)

    return router
}
