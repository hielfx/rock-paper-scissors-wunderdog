package domain

// GameStatus is the game status enum
type GameStatus string

const (
	GameStatusStarted  GameStatus = "STARTED"
	GameStatusFinished GameStatus = "FINISHED"
	GameStatusPaused   GameStatus = "PAUSED"
)
