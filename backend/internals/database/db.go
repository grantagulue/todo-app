package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)


type DB struct {
	Pool *sql.DB
}


func New() *DB {
	
	pool, err := sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	_, err = pool.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	log.Println("Database connection pool initialized successfully")

	return &DB{
		Pool: pool,
	}
}
