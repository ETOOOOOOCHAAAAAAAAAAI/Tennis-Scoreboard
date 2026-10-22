package repository

import (
	"Tennis-ScoreboardRe/internal/model"
	"database/sql"
	"fmt"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (m *MatchRepository) SaveMatch(matches model.Matches) error {
	query := "INSERT INTO Matches (Player1_ID, Player2_ID, Winner_ID) VALUES ($1, $2, $3)"
	_, err := m.db.Exec(query, matches.Player1.ID, matches.Player2.ID, matches.Winner.ID)
	if err != nil {
		return fmt.Errorf("failed to save finished match: %w", err)
	}
	return nil
}

func (m *MatchRepository) GetFinishedMatches(limit, offset int) ([]model.Matches, error) {
	query := `
		SELECT 
			m.ID, 
			p1.ID, p1.Name, 
			p2.ID, p2.Name, 
			w.ID, w.Name
		FROM Matches m
		JOIN Players p1 ON m.Player1_ID = p1.ID
		JOIN Players p2 ON m.Player2_ID = p2.ID
		JOIN Players w ON m.Winner_ID = w.ID
		ORDER BY m.ID DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := m.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query matches: %w", err)
	}
	defer rows.Close()
	var matches []model.Matches
	for rows.Next() {
		var match model.Matches
		err := rows.Scan(
			&match.ID,
			&match.Player1.ID, &match.Player1.Name,
			&match.Player2.ID, &match.Player2.Name,
			&match.Winner.ID, &match.Winner.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan match: %w", err)
		}
		matches = append(matches, match)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return matches, nil
}
