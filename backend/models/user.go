// models/user.go
package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (user *User) CreateUser(db *sql.DB) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)` // PostgreSQL uses $1, $2 for placeholders
	_, err := db.Exec(query, user.Email, user.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	row := db.QueryRow(query, email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
