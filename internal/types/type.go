// internal/types/types.go
package types

import "time"

type SensorData struct {
	AreaID    string    `json:"areaId"`
	SensorID  string    `json:"sensorId"`
	FlowRate  float64   `json:"flowRate"`
	Timestamp time.Time `json:"timestamp"`
}
