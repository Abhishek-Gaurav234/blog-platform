package models

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

var (
	instance *Database
	once     sync.Once
)

func GetDatabaseInstance() *Database {
	once.Do(func() {
		// Get database path from environment variable or use default
		dbPath := os.Getenv("DB_PATH")
		if dbPath == "" {
			dbPath = "./blog.db"
		}

		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(1) // SQLite works best with a single connection
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(5 * time.Minute)

		// Create posts table if it doesn't exist
		createTableSQL := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			type TEXT NOT NULL,
			author_id INTEGER NOT NULL,
			status TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`

		_, err = db.Exec(createTableSQL)
		if err != nil {
			panic(err)
		}

		log.Println("SQLite database initialized successfully")

		instance = &Database{DB: db}
	})
	return instance
}
