package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "lazytask.db"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

// Initialize creates a new SQLite database and sets up the required tables
func Initialize() (*DB, error) {
	// Get user's home directory
  db, err := GetDatabaseConnection()
	if err != nil {
		return nil, err
	}

	// Create tasks table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL DEFAULT 'pending',
			due_date DATETIME,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Printf("Error creating tasks table: %v", err)
		return nil, err
	}

	return db, nil
}

func GetDatabaseConnection() (*DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Create .lazytask directory if it doesn't exist
	configDir := filepath.Join(homeDir, ".lazytask")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// Open SQLite database
	dbPath := filepath.Join(configDir, dbName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

  return &DB{db}, err
}
