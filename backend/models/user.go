package models

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) CreateUser(db *sql.DB) error {
	query := `
		INSERT INTO users (email, password, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id, created_at, updated_at`
	
	err := db.QueryRow(query, user.Email, user.Password).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	row := db.QueryRow(query, email)
	
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByID(db *sql.DB, userID int) (*User, error) {
	query := `SELECT id, email, created_at, updated_at FROM users WHERE id = $1`
	row := db.QueryRow(query, userID)
	
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (user *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}