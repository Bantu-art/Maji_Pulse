package sensor

import (
    "time"
    "math/rand"
)

// Constants for cost calculation
const (
    CostPerLitrePerMinute = 0.05 // Example cost per litre per minute
    UpdateInterval        = 1 * time.Minute
)

// SensorData holds data for a sensor
type SensorData struct {
    FlowRate float64
    Cost     float64
}

// GetSensorData retrieves current sensor data
func GetSensorData() SensorData {
    // Simulate sensor data
    flowRate := rand.Float64() * 10 // Random flow rate between 0 and 10 litres per minute
    return SensorData{
        FlowRate: flowRate,
        Cost:     calculateCost(flowRate),
    }
}

// Calculate cost based on flow rate
func calculateCost(flowRate float64) float64 {
    return flowRate * CostPerLitrePerMinute
}

// Simulate sensor data over time
func StartSensorSimulation() {
    // Dummy implementation to simulate flow rate changes
    go func() {
        ticker := time.NewTicker(UpdateInterval)
        for range ticker.C {
            flowRate := rand.Float64() * 10
            cost := calculateCost(flowRate)
            // Update cost in some persistent storage
            // e.g., save to database or file
            // For demonstration, we just print it
            println("Flow Rate:", flowRate, "Cost:", cost)
        }
    }()
}

type AreaSensorData struct {
    Area     string
    FlowRate float64
    Volume   float64
}

// LeakSensorData holds data specifically for leak detection
type LeakSensorData struct {
    Area     string
    FlowRate float64
}

// SimulateAreaSensorData simulates sensor data for multiple areas
func SimulateAreaSensorData(area string) AreaSensorData {
    rand.Seed(time.Now().UnixNano())
    return AreaSensorData{
        Area:     area,
        FlowRate: rand.Float64() * 20, // Larger volume for areas
        Volume:   rand.Float64() * 1000, // Larger volume
    }
}

// SimulateLeakSensorData simulates sensor data for leak detection
func SimulateLeakSensorData(area string) LeakSensorData {
    rand.Seed(time.Now().UnixNano())
    // Randomly generate very high or very low flow rates to simulate leaks
    if rand.Intn(2) == 0 {
        return LeakSensorData{
            Area:     area,
            FlowRate: rand.Float64() * 5, // Very low flow rate
        }
    }
    return LeakSensorData{
        Area:     area,
        FlowRate: 20 + rand.Float64()*10, // Very high flow rate
    }
}
