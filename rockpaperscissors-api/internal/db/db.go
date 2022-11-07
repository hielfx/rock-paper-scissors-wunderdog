package db

import (
	"rockpaperscissors-api/internal/domain"
	"rockpaperscissors-api/internal/errors"
	"sync"

	gwebsocket "github.com/gorilla/websocket"
)

// GamesDB is a shorthand type of the "Game table"
type GamesDB map[string]domain.Game

// PlayersDB is a shorthand type for the "Player table"
type PlayersDB map[string]domain.Player

// DB holds our in memory database info, emulating tables from a real db
type DB struct {
	mu      *sync.RWMutex
	games   GamesDB
	players PlayersDB
}

// New creates and returns a new DB
func New() DB {
	return DB{
		games:   make(GamesDB),
		players: make(PlayersDB),
		mu:      &sync.RWMutex{},
	}
}

// GetPlayerByNickname Retrieves a player by it's nickname from the database
func (db *DB) GetPlayerByNickname(nick string) (player domain.Player, isOK bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	player, isOK = db.players[nick]

	return
}

// SetPlayer inserts a player into the database
func (db *DB) SetPlayer(player domain.Player) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.players[player.Nickname] = player
}

// SetGame insert's a game into the database
func (db *DB) SetGame(game domain.Game) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.games[game.ID] = game
}

// JoinGame inserts a player into the db and the game
// FIXME: Separate the inner checks into different functions
func (db *DB) JoinGame(gameID string, nickname string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	//check if game exists
	game, isOK := db.games[gameID]
	if !isOK {
		return errors.ErrGameNotFound
	}

	//check if there's already 2 players in the game
	if len(game.Players) >= 2 {
		return errors.ErrMaxPlayersReached
	}

	//check if the player is already in the room
	// alreadyJoined := false
	// for i := 0; i < len(game.Players); i++ {
	// 	if game.Players[0].Nickname == nickname {
	// 		alreadyJoined = true
	// 		break
	// 	}
	// }
	if _, alreadyJoined := game.Players[nickname]; alreadyJoined {
		return errors.ErrPlayerAlreadyJoined
	}

	//insert the player
	game.Players[nickname] = &domain.Player{Nickname: nickname}

	//Check if all the players joined
	if len(game.Players) == 2 {
		game.Status = domain.GameStatusStarted
	}

	db.games[gameID] = game

	return nil
}

// GetGameByID retrieves a game by its ID from the database
func (db *DB) GetGameByID(id string) (game domain.Game, isOK bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	game, isOK = db.games[id]

	return
}

func (db *DB) Connect(payload domain.WebsocketPayload, conn *gwebsocket.Conn) {
	db.mu.Lock()
	defer db.mu.Unlock()
	game, isOK := db.games[payload.GameID]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrGameNotFound.Error(),
		})
	}

	player, isOK := db.games[payload.GameID].Players[payload.Nickname]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrPlayerNotJoined.Error(),
		})
	}

	player.Conn = conn

	conn.WriteJSON(domain.WebsocketPayload{
		Command: domain.WebsocketPayloadCommandOK,
		Message: "Connected",
		Game:    game,
	})
}

func (db *DB) RemovePlayer(payload domain.WebsocketPayload, conn *gwebsocket.Conn) {
	db.mu.Lock()
	defer db.mu.Unlock()
	game, isOK := db.games[payload.GameID]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrGameNotFound.Error(),
		})
		return
	}

	delete(game.Players, payload.Nickname)
}

// RemoveConnection - removes the connection from the games.
// It also returns a list of games from which the connection has been removed
// and sets the game status as FINISHED
func (db *DB) RemoveConnection(conn *gwebsocket.Conn) []domain.Game {
	db.mu.Lock()
	defer db.mu.Unlock()

	gamesConnectionRemoved := []domain.Game{}

	for gameId, game := range db.games {
		for playerNick, player := range game.Players {
			if player.Conn == conn {
				player.Conn = nil
				// delete(db.games[gameId].Players, playerNick)
				game.Players[playerNick] = player
				//TODO: Apply a leaver penalty and declare the other player as winner
				game.Status = domain.GameStatusFinished
				db.games[gameId] = game
				gamesConnectionRemoved = append(gamesConnectionRemoved, db.games[gameId])
			}
		}
	}

	return gamesConnectionRemoved
}

func (db *DB) Play(payload domain.WebsocketPayload, conn *gwebsocket.Conn) *domain.Game {
	db.mu.Lock()
	defer db.mu.Unlock()

	game, isOK := db.games[payload.GameID]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrGameNotFound.Error(),
		})
		return nil
	}

	round := game.Rounds[game.CurrentRound]

	if round.PlayerSelections == nil {
		round.PlayerSelections = make(map[string]domain.Move)
	}
	if _, exists := round.PlayerSelections[payload.Nickname]; !exists {
		round.PlayerSelections[payload.Nickname] = payload.Value
	}

	//Check if all players already choose and select a winner
	//TODO: Export this to another function
	if len(round.PlayerSelections) == 2 {
		var currentWinner string
		var currentMove domain.Move
		isDraw := false
		//check if both players already choose and set the winner
		for nickname, move := range round.PlayerSelections {
			if currentWinner == "" {
				currentWinner = nickname
				currentMove = move
			} else {
				if move.Beats(currentMove) {
					currentWinner = nickname
					currentMove = move
				} else if currentMove.Beats(move) {
					//DO nothing here
				} else {
					currentWinner = ""
					isDraw = true
				}
			}
		}
		round.Winner = currentWinner
		round.IsDraw = isDraw
	}

	game.Rounds[game.CurrentRound] = round

	if round.Winner != "" || round.IsDraw {
		if game.CurrentRound < game.RoundsNumber-1 {
			game.CurrentRound++
			//Create a new round
		} else {
			game.Status = domain.GameStatusFinished

			//Declare a game winner
			currentPlayerWins := 0
			draws := 0
			for _, round := range game.Rounds {
				if round.Winner == payload.Nickname {
					currentPlayerWins++
				} else if round.IsDraw {
					draws++
				}
			}
			opponentWins := game.RoundsNumber - draws - currentPlayerWins

			if currentPlayerWins > opponentWins {
				game.Winner = payload.Nickname
			} else if opponentWins > currentPlayerWins {
				var opponent string
				for nickname := range game.Players {
					if nickname != payload.Nickname {
						opponent = nickname
						break
					}
				}
				game.Winner = opponent
			} else {
				game.IsDraw = true
			}
		}
	}

	db.games[payload.GameID] = game

	return &game

}

// Pause - pauses the game
func (db *DB) Pause(payload domain.WebsocketPayload, conn *gwebsocket.Conn) *domain.Game {
	db.mu.Lock()
	defer db.mu.Unlock()
	game, isOK := db.games[payload.GameID]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrGameNotFound.Error(),
		})
		return nil
	}

	game.Status = domain.GameStatusPaused
	game.PausedBy = payload.Nickname

	db.games[payload.GameID] = game

	return &game
}

// Unpause - unpauses the game
func (db *DB) Unpause(payload domain.WebsocketPayload, conn *gwebsocket.Conn) *domain.Game {
	db.mu.Lock()
	defer db.mu.Unlock()
	game, isOK := db.games[payload.GameID]
	if !isOK {
		conn.WriteJSON(domain.WebsocketPayload{
			Command: domain.WebsocketPayloadCommandError,
			Message: errors.ErrGameNotFound.Error(),
		})
		return nil
	}

	//Check if it's the same player that paused the game
	if payload.Nickname == game.PausedBy {
		game.Status = domain.GameStatusStarted
		game.PausedBy = ""
	}

	db.games[payload.GameID] = game

	return &game
}
