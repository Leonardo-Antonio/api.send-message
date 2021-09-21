package app

import (
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/api.send-message/src/env"
	"github.com/Leonardo-Antonio/api.send-message/src/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	app *echo.Echo
}

func NewApp() *server {
	return &server{
		app: echo.New(),
	}
}

func (s *server) Middlewares() {
	s.app.Use(middleware.CORS())
	s.app.Use(middleware.Logger())
}

func (s *server) Routers() {
	router.Message(s.app)
}

func (s *server) Listeing() {
	log.Fatalln(s.app.Start(fmt.Sprintf(":%s", env.Port)))
}
