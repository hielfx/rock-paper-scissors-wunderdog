package routes

import (
	"rockpaperscissors-api/internal/handler"

	"github.com/labstack/echo/v4"
)

func AppendGameRoutes(e *echo.Group, h handler.Handler) {
	e.POST("", h.CreateGame)
	// e.GET("", h.GetGames)
	e.POST("/:gameId", h.JoinGame)
}
