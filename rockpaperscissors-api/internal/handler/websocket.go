package handler

import (
	"net/http"
	"rockpaperscissors-api/internal/domain"
	"rockpaperscissors-api/internal/errors"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: checkOrigins,
	}
)

func (h *Handler) Websocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		logrus.Errorf("Error in handler.Websocket -> error upgrading the connection: %s", err)
		return err
	}
	defer ws.Close()
	ws.SetCloseHandler(func(code int, text string) error {
		ws.Close()
		return nil
	})

	for {

		var payload domain.WebsocketPayload
		if err := ws.ReadJSON(&payload); err != nil {
			//revemo the connection if there's an error
			// games := h.db.RemoveConnection(ws)
			// for _, g := range games {
			// 	sendGameInfoToAllPlayers(g)
			// }

			// Only log "unexpected" errors
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				games := h.db.RemoveConnection(ws)
				for _, g := range games {
					sendGameInfoToAllPlayers(g)
				}
			} else {
				logrus.Errorf("Error in handler.Websocket -> error reading json payload: %s", err)

			}

			return err
		}

		//TODO: check if payload is valid

		switch payload.Command {
		case domain.WebsocketPayloadCommandConnect:
			h.db.Connect(payload, ws)
			game, isOK := h.db.GetGameByID(payload.GameID)
			if !isOK {
				return echo.NewHTTPError(http.StatusNotFound, errors.ErrMissingGame)
			}
			sendGameInfoToAllPlayers(game)
		case domain.WebsocketPayloadCommandClose:
			if err := ws.Close(); err != nil {
				logrus.Errorf("Error in handler.Websocket -> error closing the connection: %s", err)
				return err
			}
			//Only remove the player if the game hasn't finished
			// game, isOK := h.db.GetGameByID(payload.GameID)
			// if !isOK {
			// 	return echo.NewHTTPError(http.StatusNotFound, errors.ErrMissingGame)
			// }
			// if game.Status != domain.GameStatusFinished {
			// 	h.db.RemovePlayer(payload, ws)
			// 	game, isOK := h.db.GetGameByID(payload.GameID)
			// 	if !isOK {
			// 		return echo.NewHTTPError(http.StatusNotFound, errors.ErrMissingGame)
			// 	}
			// 	sendGameInfoToAllPlayers(game)
			// }
			games := h.db.RemoveConnection(ws)
			for _, g := range games {
				sendGameInfoToAllPlayers(g)
			}
		case domain.WebsocketPayloadCommandPlay:
			game := h.db.Play(payload, ws)
			// game, isOK := h.db.GetGameByID(payload.GameID)
			// if !isOK {
			// 	return echo.NewHTTPError(http.StatusNotFound, errors.ErrMissingGame)
			// }
			if game != nil {
				sendGameInfoToAllPlayers(*game)
			}
		case domain.WebsocketPayloadCommandPause:
			game := h.db.Pause(payload, ws)
			if game != nil {
				sendGameInfoToAllPlayers(*game)
			}
		case domain.WebsocketPayloadCommandUnpause:
			game := h.db.Unpause(payload, ws)
			if game != nil {
				sendGameInfoToAllPlayers(*game)
			}
		default:
			if err := ws.WriteJSON(domain.WebsocketPayload{
				Command: domain.WebsocketPayloadCommandError,
				Message: errors.ErrInvalidCommand.Error(),
			}); err != nil {
				logrus.Errorf("Error in handler.Websocket -> error writing json: %s", err)
				return err
			}
		}
	}
}

func sendGameInfoToAllPlayers(game domain.Game) {
	for _, p := range game.Players {
		if p != nil {
			p.Conn.WriteJSON(domain.WebsocketPayload{
				Command:  domain.WebsocketPayloadCommandOK,
				GameID:   game.ID,
				Game:     game,
				Nickname: p.Nickname,
			})
		}
	}
}

// import (
// 	"fmt"
// 	_websocket "rockpaperscissors-api/internal/websocket"

// 	"github.com/labstack/echo/v4"
// 	"github.com/sirupsen/logrus"
// )

// func (h *Handler) Websocket(c echo.Context) error {
// 	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		logrus.Errorf("Error in routes.Websocket -> error upgrading conection: %s", err)
// 		return err
// 	}
// 	defer ws.Close()

// 	client := &_websocket.Client{
// 		Hub:  h.hub,
// 		Conn: ws,
// 	}

// 	h.hub.Register <- client

// 	for {

// 		mt, msg, err := ws.ReadMessage()
// 		if err != nil {
// 			logrus.Errorf("Error in routes.Websocket -> error reading message: %s", err)
// 			return err
// 		}

// 		fmt.Println(string(msg))

// 		if err := ws.WriteMessage(mt, msg); err != nil {
// 			logrus.Errorf("Error in routes.Websocket -> error writting message: %s", err)
// 			return err
// 		}
// 	}

// 	return nil
// }
