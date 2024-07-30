package config

import (
	"os"
)

// Config holds the configuration settings for the application.
type Config struct {
	ServerPort    string // Port for the API server
	BlockchainURL string // URL for the blockchain service
	SensorAPIURL  string // API endpoint for sensor data
}

// LoadConfig loads the application configuration from environment variables or defaults.
func LoadConfig() Config {
	return Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		BlockchainURL: getEnv("BLOCKCHAIN_URL", "http://localhost:8000"),
		SensorAPIURL:  getEnv("SENSOR_API_URL", "http://localhost:8080/api/sensor-data"),
	}
}

// Helper function to get environment variable as string
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper function to get environment variable as int
// func getEnvAsInt(key string, fallback int) int {
// 	if value, exists := os.LookupEnv(key); exists {
// 		if intValue, err := strconv.Atoi(value); err == nil {
// 			return intValue
// 		}
// 	}
// 	return fallback
// }
