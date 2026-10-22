package service

import (
	"Tennis-ScoreboardRe/internal/model"
	"Tennis-ScoreboardRe/internal/repository"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type MatchService struct {
	playerRepo             *repository.PlayerRepository
	ongMatchRepo           *repository.OngMatchRepository
	matchRepo              *repository.MatchRepository
	scoreCalculatorService *MatchScoreCalculationService
}

func NewMatchService(pl *repository.PlayerRepository, ong *repository.OngMatchRepository, mr *repository.MatchRepository, ss *MatchScoreCalculationService) *MatchService {
	return &MatchService{playerRepo: pl,
		ongMatchRepo:           ong,
		matchRepo:              mr,
		scoreCalculatorService: ss,
	}
}

func (s *MatchService) GetOrCreatePlayer(name string) (model.Players, error) {
	player, err := s.playerRepo.GetPlayerByName(name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // расскрытие матрешки ошибок
			return s.playerRepo.CreatePlayer(name)
		}
		return model.Players{}, err
	}
	return player, nil
}

func (s *MatchService) CreateMatch(player1Name, player2Name string) (uuid.UUID, error) {
	pl1, err := s.GetOrCreatePlayer(player1Name)
	if err != nil {
		return uuid.Nil, err
	}
	pl2, err := s.GetOrCreatePlayer(player2Name)
	if err != nil {
		return uuid.Nil, err
	}
	matchID := uuid.New()
	match := model.Matches{Player1: pl1,
		Player2: pl2}
	s.ongMatchRepo.Save(matchID, match)
	return matchID, nil
}

func (s *MatchService) ProcessPoint(matchID uuid.UUID, winnerID int64) error {
	var winnerNumber int
	match, exists := s.ongMatchRepo.Get(matchID)
	if !exists {
		return errors.New("match not found")
	}
	if winnerID == match.Player1.ID {
		winnerNumber = 1
	} else if winnerID == match.Player2.ID {
		winnerNumber = 2
	}
	s.scoreCalculatorService.AddPoint(&match.Score, winnerNumber)
	s.scoreCalculatorService.CheckGameWinner(&match.Score)
	if s.scoreCalculatorService.IsMatchFinished(&match.Score) {
		if match.Score.Player1Sets == 2 {
			match.Winner = match.Player1
		} else if match.Score.Player2Sets == 2 {
			match.Winner = match.Player2
		}
		s.ongMatchRepo.Save(matchID, match)
		return s.FinishMatch(matchID)

	}
	s.ongMatchRepo.Save(matchID, match)
	return nil

}

func (s *MatchService) FinishMatch(ID uuid.UUID) error {
	match, exists := s.ongMatchRepo.Get(ID)
	if !exists {
		return errors.New("match not found")
	}
	err := s.matchRepo.SaveMatch(match)
	if err != nil {
		return err
	}
	s.ongMatchRepo.Delete(ID)
	return nil
}

func (s *MatchService) GetMatch(ID uuid.UUID) (model.Matches, error) {
	if match, exists := s.ongMatchRepo.Get(ID); exists {
		return match, nil
	}
	return model.Matches{}, errors.New("Match not found")
}

func (s *MatchService) GetFinishedMatch(page int) ([]model.Matches, error) {
	limit := 5
	offset := (page - 1) * limit
	return s.matchRepo.GetFinishedMatches(limit, offset)
}
