// models/journal_entry.go
package models

import (
	"database/sql"
	"log"
)

type JournalEntry struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func (entry *JournalEntry) CreateEntry(db *sql.DB) error {
	query := `INSERT INTO journals (content, user_id) VALUES ($1, $2)`
	_, err := db.Exec(query, entry.Content, entry.UserID)
	log.Println("Error creating journal entry:", err)
	if err != nil {
		log.Printf("Error 1 creating journal entry: %v (content: %s, user_id: %d)\n", err, entry.Content, entry.UserID)
		return err
	}
	return nil
}

func GetEntriesByUser(db *sql.DB, userID int) ([]JournalEntry, error) {
	query := `SELECT id, content, user_id FROM journals WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []JournalEntry
	for rows.Next() {
		var entry JournalEntry
		err := rows.Scan(&entry.ID, &entry.Content, &entry.UserID)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
