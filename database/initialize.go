package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "jnda.db"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

var database *DB

// Initialize creates a new SQLite database and sets up the required tables
func init_db(db *DB) error {
  var err error

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
		return err
	}

	return nil
}

func GetDatabaseConnection() (*DB, error) {
  if database != nil {
    return database, nil
  }

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Create .jnda directory if it doesn't exist
	configDir := filepath.Join(homeDir, ".jnda")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	// Open SQLite database
	dbPath := filepath.Join(configDir, dbName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
  database = &DB{db} 

  res, err := database.Query(`
    SELECT name 
    FROM sqlite_master 
    WHERE type='table' AND name='tasks';
  `)
  if err != nil || res.Next() == false  {
    log.Print("Init db")
    err = init_db(database)
    if err != nil {
      log.Fatal("Failed to init database")
    }
  }
  defer res.Close()

	return database, err
}
