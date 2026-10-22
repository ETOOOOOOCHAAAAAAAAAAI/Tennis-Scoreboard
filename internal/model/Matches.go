package model

type Matches struct {
	ID      int64
	Player1 Players
	Player2 Players
	Winner  Players
	Score   Score
}
