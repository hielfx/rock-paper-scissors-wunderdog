package websocket

import "rockpaperscissors-api/internal/domain"

// GameWS Game object for websocket.
// It has the game data with a map of players websocket connections (Client)
type GameWS struct {
	domain.Game
	Players map[*Client]bool `json:"players"`
}
