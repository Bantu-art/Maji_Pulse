package models

import (
	"database/sql"
	"time"
)

// Leakage represents a leakage record.
type Leakage struct {
	ID        int       `json:"id"`
	SensorID  string    `json:"sensor_id"`
	FlowRate  float64   `json:"flow_rate"`
	Timestamp time.Time `json:"timestamp"`
}

// GetLeakages retrieves leakage records from the database.
func GetLeakages(db *sql.DB) ([]Leakage, error) {
	rows, err := db.Query("SELECT id, sensor_id, flow_rate, timestamp FROM leakage")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leakages []Leakage
	for rows.Next() {
		var leakage Leakage
		if err := rows.Scan(&leakage.ID, &leakage.SensorID, &leakage.FlowRate, &leakage.Timestamp); err != nil {
			return nil, err
		}
		leakages = append(leakages, leakage)
	}
	return leakages, nil
}

// InsertLeakage inserts a new leakage record into the database.
func InsertLeakage(db *sql.DB, sensorID string, flowRate float64) error {
	stmt, err := db.Prepare("INSERT INTO leakage(sensor_id, flow_rate) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensorID, flowRate)
	return err
}
