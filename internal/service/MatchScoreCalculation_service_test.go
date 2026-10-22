package service

import (
	"Tennis-ScoreboardRe/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPoint_Player1ScoreFirstPoint(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.Point0, Player2Point: model.Point0}
	calc.AddPoint(score, 1)
	assert.Equal(t, model.Point15, score.Player1Point, "Очки первого игрока должны стать 15")
	assert.Equal(t, model.Point0, score.Player2Point, "Очки второго игрока не должны были измениться")
	assert.Equal(t, 0, score.Player1Game, "Количество геймов не должно было вырасти")
}

func TestAddPoint_Player2ScoreFirstPoint(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.Point0, Player2Point: model.Point0}
	calc.AddPoint(score, 2)
	assert.Equal(t, model.Point15, score.Player2Point, "Очки первого игрока должны стать 15")
	assert.Equal(t, model.Point0, score.Player1Point, "Очки второго игрока не должны были измениться")
	assert.Equal(t, 0, score.Player1Game, "Количество геймов не должно было вырасти")
}

func TestAddPoint_CheckTieBreakPlayer2(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{
		Player1Point: model.Point0, Player2Point: model.Point40,
		Player1Game: 6, Player2Game: 5}
	calc.AddPoint(score, 2)
	calc.CheckGameWinner(score)
	assert.Equal(t, true, score.IsTieBreak)
	assert.Equal(t, 0, score.Player1TieBreakPoint)
	assert.Equal(t, 0, score.Player2TieBreakPoint)
}

func TestAddPoint_CheckTieBreakPlayer1(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{
		Player1Point: model.Point40, Player2Point: model.Point0,
		Player1Game: 5, Player2Game: 6}
	calc.AddPoint(score, 1)
	calc.CheckGameWinner(score)
	assert.Equal(t, true, score.IsTieBreak)
	assert.Equal(t, 0, score.Player1TieBreakPoint)
	assert.Equal(t, 0, score.Player2TieBreakPoint)
}

func TestAddPoint_CheckADPlayer1Win(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.PointAD, Player2Point: model.Point40}
	calc.AddPoint(score, 1)
	assert.Equal(t, model.Point0, score.Player1Point)
	assert.Equal(t, model.Point0, score.Player2Point)
	assert.Equal(t, 1, score.Player1Game)
	assert.Equal(t, 0, score.Player2Game)
}

func TestAddPoint_CheckADPlayer2Win(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.Point40, Player2Point: model.PointAD}
	calc.AddPoint(score, 2)
	assert.Equal(t, model.Point0, score.Player1Point)
	assert.Equal(t, model.Point0, score.Player2Point)
	assert.Equal(t, 0, score.Player1Game)
	assert.Equal(t, 1, score.Player2Game)
}

func TestAddPoint_CheckADPlayer2LoseAD(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.Point40, Player2Point: model.PointAD}
	calc.AddPoint(score, 1)
	assert.Equal(t, model.Point40, score.Player1Point) // проблема была не в том что я неправильно нааписал логику а просто сам запутался в правилах тенниса
	assert.Equal(t, model.Point40, score.Player2Point)

}

func TestAddPoint_CheckADPlayer1LoseAD(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{Player1Point: model.PointAD, Player2Point: model.Point40}
	calc.AddPoint(score, 2)
	assert.Equal(t, model.Point40, score.Player1Point)
	assert.Equal(t, model.Point40, score.Player2Point)

}

func TestMatchFinishing(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{
		Player1Point: model.Point40,
		Player2Point: model.Point0,
		Player1Game:  6,
		Player2Game:  0,
		Player1Sets:  1,
		Player2Sets:  0}
	calc.AddPoint(score, 1)
	calc.CheckGameWinner(score)
	assert.Equal(t, true, calc.IsMatchFinished(score))

}

func TestMatchFinishingPlayer2(t *testing.T) {
	calc := NewMatchScoreCalculationService()
	score := &model.Score{
		Player1Point: model.Point15,
		Player2Point: model.Point40,
		Player1Game:  0,
		Player2Game:  6,
		Player1Sets:  1,
		Player2Sets:  1}
	calc.AddPoint(score, 2)
	calc.CheckGameWinner(score)
	assert.Equal(t, true, calc.IsMatchFinished(score))

}
