package app

import (
	"net/http"
	"os"
	"rockpaperscissors-api/internal/handler"
	"rockpaperscissors-api/internal/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	echoLogrusMiddleware "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

type App struct {
	router  *echo.Echo
	handler handler.Handler
}

type AppConfig struct {
	Debug bool
}

var DefaultConfig = AppConfig{
	Debug: true,
}

func New() App {
	return NewWithConfig(DefaultConfig)
}

func NewWithConfig(config AppConfig) App {
	e := echo.New()

	e.Debug = config.Debug

	setupMiddlewares(e)

	h := handler.NewHandler()

	app := App{
		router:  e,
		handler: h,
	}

	app.initializehandler()

	return app
}

func (a *App) Start(address string) error {
	return a.router.Start(address)
}

// func (a *App) InitializeHub() {
// 	a.handler.RunHub()
// }

func setupMiddlewares(e *echo.Echo) {
	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	e.Use(echoLogrusMiddleware.Logger())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, "Cache-Control"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
}

func (a *App) initializehandler() {
	a.router.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "Alive!",
		})
	})

	a.router.GET("/ws", a.handler.Websocket)

	routes.AppendGameRoutes(a.router.Group("/games"), a.handler)

}
