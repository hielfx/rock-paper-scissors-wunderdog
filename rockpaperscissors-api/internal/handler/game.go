package handler

import (
	"net/http"
	"rockpaperscissors-api/internal/domain"
	"rockpaperscissors-api/internal/errors"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CreateGameBody struct {
	Rounds int `json:"rounds"`
}

// CreateGame godoc
// @Description Creates and stores a game in the "db"
// @Accept json
// @Produce json
// @Success 201 {object} domain.Game
// @Failure 400 {object} echo.HttpError
// @Failure 500 {object} error
// @Router /games [post]
func (h *Handler) CreateGame(c echo.Context) error {
	var body CreateGameBody

	if err := c.Bind(&body); err != nil {
		logrus.Errorf("Error in handler.CreateGame -> error parsing body: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	game := domain.NewGame(body.Rounds)

	h.db.SetGame(game)

	// h.hub.NewGame(game)

	return c.JSON(http.StatusCreated, game)
}

type joinGameBody struct {
	Nickname string `json:"nickname"`
}

// JoinGame godoc
// @Description Joins a player to a game
// @Accept json
// @Produce json
// @Success 200 {object} domain.Game
// @Failure 400 {object} echo.HttpError
// @Failure 500 {object} error
// @Router /games/:gameId [post]
func (h *Handler) JoinGame(c echo.Context) error {
	gameID := c.Param("gameId")
	if gameID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.ErrInvalidGameID)
	}

	var body joinGameBody
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//join the game
	if err := h.db.JoinGame(gameID, body.Nickname); err != nil {
		return joinGameErrorGenerator(err)
	}

	res, isOK := h.db.GetGameByID(gameID)
	if !isOK {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.ErrMissingGame)
	}
	return c.JSON(http.StatusOK, res)
}

// func (h *Handler) JoinGame(c echo.Context) error {
// 	ws, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		logrus.Errorf("Error in handler.JoinGame -> error upgrading conection: %s", err)
// 		return err
// 	}
// 	defer ws.Close()

// 	gameID := uuid.FromString(c.Param("gameId"))
// 	playerNick := c.QueryParam("nick")
// 	if playerNick == "" {
// 		return c.JSON(http.StatusBadRequest, errors.New("invalidPlayerNick"))
// 	}

// 	// player, isOK := h.db.GetPlayer(playerNick)
// 	// if !isOK {
// 	// 	player.Nick = playerNick
// 	// }

// 	player := websocket.Player{
// 		Nick:    playerNick,
// 		Conn:    ws,
// 		Message: make(chan *domain.Message),
// 	}

// 	h.hub.Join <- &websocket.JoinMessage{
// 		GameID: gameID,
// 		Player: player,
// 	}

// 	for {
// 		_, msg, err := ws.ReadMessage()
// 		if err != nil {
// 			logrus.Errorf("Error in handler.JoinGame")
// 		}

// 		fmt.Print(msg)

// 	}

// 	return nil
// }

func joinGameErrorGenerator(err error) *echo.HTTPError {
	var code int

	switch err {
	case errors.ErrGameNotFound:
		code = http.StatusNotFound
	case errors.ErrMaxPlayersReached, errors.ErrPlayerAlreadyJoined:
		code = http.StatusConflict
	case errors.ErrInvalidGameID:
		fallthrough
	default:
		code = http.StatusBadRequest
	}

	return echo.NewHTTPError(code, err.Error())
}
