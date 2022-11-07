package domain

import (
	"errors"

	"github.com/google/uuid"
)

type MessageType string

const (
	MessageTypeError  MessageType = "ERROR"
	MessageTypeCreate MessageType = "CREATE"
	MessageTypeJOIN   MessageType = "JOIN"
	MessageTypePAUSE  MessageType = "PAUSE"
)

type GameError error

var (
	GameErrorGameNotFound GameError = errors.New("gameNotFound")
)

type Message struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
	From    *string     `json:"from"`
	GameID  *uuid.UUID  `json:"gameId"`
	Error   GameError
}
