package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database, %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to real open database, %w", err)
	}
	return db, nil
}
