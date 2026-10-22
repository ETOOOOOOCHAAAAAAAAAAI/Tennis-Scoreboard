package repository

import (
	"Tennis-ScoreboardRe/internal/model"
	"database/sql"
	"fmt"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) GetPlayerByName(name string) (model.Players, error) {
	var p model.Players
	err := r.db.QueryRow("SELECT ID, Name FROM PLAYERS WHERE Name = $1", name).Scan(&p.ID, &p.Name)
	if err != nil {
		return p, fmt.Errorf("Error requesting player name: %w", err)
	}
	return p, nil
}

func (r *PlayerRepository) CreatePlayer(name string) (model.Players, error) {
	var p model.Players
	err := r.db.QueryRow("INSERT INTO Players (Name) VALUES ($1) RETURNING ID", name).Scan(&p.ID)
	if err != nil {
		return p, fmt.Errorf("Player already exists: %w", err)
	}
	p.Name = name
	return p, nil
}
