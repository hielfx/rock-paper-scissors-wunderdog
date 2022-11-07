package domain

import (
	"github.com/google/uuid"
	gwebsocket "github.com/gorilla/websocket"
)

// Game object
type Game struct {
	ID           string             `json:"id"`
	RoundsNumber int                `json:"roundsNumber"`
	Players      map[string]*Player `json:"players"`
	Winner       string             `json:"winner"`
	Rounds       []Round            `json:"rounds"`
	Status       GameStatus         `json:"status"`
	CurrentRound int                `json:"currentRound"`
	IsDraw       bool               `json:"isDraw"`
	PausedBy     string             `json:"pausedBy"`
}

// NewGame creates a new game with the given rounds and a generated ID
func NewGame(roundsNumber int) Game {
	if roundsNumber < 1 {
		roundsNumber = 1
	}
	return Game{
		ID:           uuid.NewString(),
		RoundsNumber: roundsNumber,
		Rounds:       make([]Round, roundsNumber),
		Players:      make(map[string]*Player),
		CurrentRound: 0,
	}
}

// Player object
type Player struct {
	Nickname string           `json:"nickname"`
	Conn     *gwebsocket.Conn `json:"-"`
}

// Round players selections and winner for a single round
type Round struct {
	PlayerSelections map[string]Move `json:"playerSelections"`
	Winner           string          `json:"winner"`
	IsDraw           bool            `json:"isDraw"`
}

// Selection move selection by a single player
type Selection struct {
	MadeBy Player
	Move   Move
}
