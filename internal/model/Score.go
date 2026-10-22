package model

type PlayerPoint int

const (
	Point0 PlayerPoint = iota
	Point15
	Point30
	Point40
	PointAD
)

type Score struct {
	Player1Sets          int
	Player2Sets          int
	Player1Game          int
	Player2Game          int
	Player1Point         PlayerPoint
	Player2Point         PlayerPoint
	IsTieBreak           bool
	Player1TieBreakPoint int
	Player2TieBreakPoint int
}

func (p PlayerPoint) String() string {
	switch p {
	case Point0:
		return "0"
	case Point15:
		return "15"
	case Point30:
		return "30"
	case Point40:
		return "40"
	case PointAD:
		return "AD"
	default:
		return "0"
	}
}
