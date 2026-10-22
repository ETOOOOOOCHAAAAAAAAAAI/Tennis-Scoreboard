package service

import (
	"Tennis-ScoreboardRe/internal/model"
)

type MatchScoreCalculationService struct{}

func NewMatchScoreCalculationService() *MatchScoreCalculationService {
	return &MatchScoreCalculationService{}
}

func (m *MatchScoreCalculationService) AddPoint(score *model.Score, winnerPlayerNumber int) {
	if score.IsTieBreak {
		if winnerPlayerNumber == 1 {
			score.Player1TieBreakPoint++
			if score.Player1TieBreakPoint >= 7 && score.Player1TieBreakPoint-score.Player2TieBreakPoint >= 2 {
				score.Player1TieBreakPoint = 0
				score.Player2TieBreakPoint = 0
				score.Player1Game = 0
				score.Player2Game = 0
				score.Player1Sets++
				score.IsTieBreak = false
			}
		} else if winnerPlayerNumber == 2 {
			score.Player2TieBreakPoint++
			if score.Player2TieBreakPoint >= 7 && score.Player2TieBreakPoint-score.Player1TieBreakPoint >= 2 {
				score.Player1TieBreakPoint = 0
				score.Player2TieBreakPoint = 0
				score.Player1Game = 0
				score.Player2Game = 0
				score.Player2Sets++
				score.IsTieBreak = false
			}

		}
	} else {
		if winnerPlayerNumber == 1 {
			if score.Player2Point == model.PointAD {
				score.Player2Point = model.Point40
				return
			}
			switch score.Player1Point {
			case model.Point0:
				score.Player1Point = model.Point15
			case model.Point15:
				score.Player1Point = model.Point30
			case model.Point30:
				score.Player1Point = model.Point40
			case model.Point40:
				if score.Player2Point == model.Point40 {
					score.Player1Point = model.PointAD
				} else {
					score.Player1Point = model.Point0
					score.Player2Point = model.Point0
					score.Player1Game++
				}
			case model.PointAD:
				score.Player1Point = model.Point0
				score.Player2Point = model.Point0
				score.Player1Game++
			}
		} else if winnerPlayerNumber == 2 {
			if score.Player1Point == model.PointAD {
				score.Player1Point = model.Point40
				return
			}
			switch score.Player2Point {
			case model.Point0:
				score.Player2Point = model.Point15
			case model.Point15:
				score.Player2Point = model.Point30
			case model.Point30:
				score.Player2Point = model.Point40
			case model.Point40:
				if score.Player1Point == model.Point40 {
					score.Player2Point = model.PointAD
				} else {
					score.Player1Point = model.Point0
					score.Player2Point = model.Point0
					score.Player2Game++
				}
			case model.PointAD:
				score.Player1Point = model.Point0
				score.Player2Point = model.Point0
				score.Player2Game++
			}
		}
	}
}

func (m *MatchScoreCalculationService) CheckGameWinner(score *model.Score) {
	if score.Player1Game >= 6 && score.Player1Game-score.Player2Game >= 2 {
		score.Player1Game = 0
		score.Player2Game = 0
		score.Player1Sets++
	} else if score.Player2Game >= 6 && score.Player2Game-score.Player1Game >= 2 {
		score.Player1Game = 0
		score.Player2Game = 0
		score.Player2Sets++
	} else if score.Player1Game == 6 && score.Player2Game == 6 {
		score.IsTieBreak = true
	}
}

func (m *MatchScoreCalculationService) IsMatchFinished(score *model.Score) bool {
	return score.Player1Sets == 2 || score.Player2Sets == 2
}
