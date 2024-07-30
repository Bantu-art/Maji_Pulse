package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var db *sql.DB

// InitDB initializes the SQLite database connection and creates necessary tables.
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./water_management.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables if they don't exist
	createTables()
}

// createTables creates the necessary tables in the database.
func createTables() {
	createWaterUsageTable := `
    CREATE TABLE IF NOT EXISTS water_usage (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id TEXT NOT NULL,
        amount REAL NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	createLeakageTable := `
    CREATE TABLE IF NOT EXISTS leakage (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sensor_id TEXT NOT NULL,
        flow_rate REAL NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	createWastageTable := `
    CREATE TABLE IF NOT EXISTS wastage (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id TEXT NOT NULL,
        description TEXT NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	// Execute the table creation statements
	if _, err := db.Exec(createWaterUsageTable); err != nil {
		log.Fatalf("Error creating water_usage table: %v", err)
	}
	if _, err := db.Exec(createLeakageTable); err != nil {
		log.Fatalf("Error creating leakage table: %v", err)
	}
	if _, err := db.Exec(createWastageTable); err != nil {
		log.Fatalf("Error creating wastage table: %v", err)
	}
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
