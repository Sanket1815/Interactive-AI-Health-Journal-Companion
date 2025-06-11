package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewConnection(connStr string) (*Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Database connected successfully")
	
	database := &Database{db}
	
	// Initialize database schema
	if err := database.initSchema(); err != nil {
		return nil, fmt.Errorf("error initializing schema: %v", err)
	}

	return database, nil
}

func (db *Database) initSchema() error {
	// Create users table
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	// Create journals table
	journalsTable := `
	CREATE TABLE IF NOT EXISTS journals (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		analysis TEXT,
		sentiment VARCHAR(20) DEFAULT 'neutral',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX IF NOT EXISTS idx_journals_user_id ON journals(user_id);
	CREATE INDEX IF NOT EXISTS idx_journals_created_at ON journals(created_at);
	`

	// Execute schema creation
	if _, err := db.Exec(usersTable); err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	if _, err := db.Exec(journalsTable); err != nil {
		return fmt.Errorf("error creating journals table: %v", err)
	}

	log.Println("Database schema initialized successfully")
	return nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

func (db *Database) HealthCheck() error {
	return db.Ping()
}