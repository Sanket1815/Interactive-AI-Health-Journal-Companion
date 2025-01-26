
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" 
)

func OpenConnection() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=1998sanket dbname=mental_health_journal sslmode=disable"

	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected")
	return db, nil
}
