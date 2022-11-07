package websocket

import (
	"rockpaperscissors-api/internal/domain"
	"sync"
)

// Hub holds all the active clients connections and games
type Hub struct {
	sync.RWMutex
	//registered game wit its players
	Games map[string]GameWS
	//register request
	register chan subscription
	//unregister request
	unregister chan subscription
	broadcast  chan []byte

	// Games      map[string]*Game
	// CreateGame chan CreateGamePayload
	// JoinGame   chan JoinGamePayload
	// LeaveGame  chan LeaveGamePayload
}

// NewGame attach a new game in the Hub, overwritting existing games with the same ID
func (h *Hub) NewGame(game domain.Game) {
	h.Lock()
	defer h.Unlock()
	h.Games[game.ID] = GameWS{
		Game:    game,
		Players: make(map[*Client]bool),
	}
}

type subscription struct {
	gameID string
	client *Client
}

func (h *Hub) Run() error {
	// for {
	// 	select {
	// 	case s := <-h.register:
	// 		game, gameExists := h.games[s.gameID]
	// 		if !gameExists {
	// 			//create a new game if the game doesn't exists (fitst connection)
	// 			game = GameWS{
	// 				Game: domain.NewGame(),,
	// 				Players: make(map[*Client]bool),
	// 			}
	// 			h.games[game.ID] = game
	// 		}
	// 		h.games[game.ID].Players[s.client] = true
	// 	case s := <-h.unregister:
	// 		game, gameExists := h.games[s.gameID]
	// 		if gameExists {
	// 			if _, playerExists := game.Players[s.client]; playerExists {
	// 				//remove the player from the game
	// 				delete(game.Players, s.client)
	// 				//close the player channel
	// 				close(s.client.send)

	// 				//if there're no players left in the game, delete the game
	// 				if len(game.Players) == 0 {
	// 					delete(h.games, s.gameID)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	return nil
}

func NewHub() *Hub {
	return &Hub{
		Games:      make(map[string]GameWS),
		register:   make(chan subscription),
		unregister: make(chan subscription),
		broadcast:  make(chan []byte),
	}
}
