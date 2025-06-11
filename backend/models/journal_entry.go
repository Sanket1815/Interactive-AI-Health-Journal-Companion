package models

import (
	"database/sql"
	"errors"
	"time"
)

type JournalEntry struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Analysis  string    `json:"analysis,omitempty"`
	Sentiment string    `json:"sentiment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JournalEntryResponse struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Analysis  string    `json:"analysis,omitempty"`
	Sentiment string    `json:"sentiment,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (entry *JournalEntry) CreateEntry(db *sql.DB) error {
	query := `
		INSERT INTO journals (content, user_id, analysis, sentiment, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, NOW(), NOW()) 
		RETURNING id, created_at, updated_at`
	
	err := db.QueryRow(query, entry.Content, entry.UserID, entry.Analysis, entry.Sentiment).Scan(
		&entry.ID, &entry.CreatedAt, &entry.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetEntriesByUser(db *sql.DB, userID int, limit, offset int) ([]JournalEntry, error) {
	query := `
		SELECT id, content, user_id, analysis, sentiment, created_at, updated_at 
		FROM journals 
		WHERE user_id = $1 
		ORDER BY created_at DESC 
		LIMIT $2 OFFSET $3`
	
	rows, err := db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []JournalEntry
	for rows.Next() {
		var entry JournalEntry
		err := rows.Scan(
			&entry.ID, &entry.Content, &entry.UserID, 
			&entry.Analysis, &entry.Sentiment, 
			&entry.CreatedAt, &entry.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func GetEntryByID(db *sql.DB, entryID, userID int) (*JournalEntry, error) {
	query := `
		SELECT id, content, user_id, analysis, sentiment, created_at, updated_at 
		FROM journals 
		WHERE id = $1 AND user_id = $2`
	
	row := db.QueryRow(query, entryID, userID)
	
	var entry JournalEntry
	err := row.Scan(
		&entry.ID, &entry.Content, &entry.UserID, 
		&entry.Analysis, &entry.Sentiment, 
		&entry.CreatedAt, &entry.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("journal entry not found")
		}
		return nil, err
	}
	return &entry, nil
}

func (entry *JournalEntry) ToResponse() JournalEntryResponse {
	return JournalEntryResponse{
		ID:        entry.ID,
		Content:   entry.Content,
		Analysis:  entry.Analysis,
		Sentiment: entry.Sentiment,
		CreatedAt: entry.CreatedAt,
	}
}