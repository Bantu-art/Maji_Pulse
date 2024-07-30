package models

import (
	"database/sql"
	"time"
)

// Wastage represents a wastage record.
type Wastage struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

// GetWastages retrieves wastage records for a specific user from the database.
func GetWastages(db *sql.DB, userID string) ([]Wastage, error) {
	rows, err := db.Query("SELECT id, user_id, description, timestamp FROM wastage WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wastages []Wastage
	for rows.Next() {
		var wastage Wastage
		if err := rows.Scan(&wastage.ID, &wastage.UserID, &wastage.Description, &wastage.Timestamp); err != nil {
			return nil, err
		}
		wastages = append(wastages, wastage)
	}
	return wastages, nil
}

// InsertWastage inserts a new wastage record into the database.
func InsertWastage(db *sql.DB, userID, description string) error {
	stmt, err := db.Prepare("INSERT INTO wastage(user_id, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, description)
	return err
}
