package domain

type Move string

const (
	MoveRock     Move = "ROCK"
	MovePaper    Move = "PAPER"
	MoveScissors Move = "SCISSORS"
)

func (m Move) Beats(move Move) bool {
	switch m {
	case MoveRock:
		return move == MoveScissors
	case MovePaper:
		return move == MoveRock
	case MoveScissors:
		return move == MovePaper
	default:
		return false
	}
}
