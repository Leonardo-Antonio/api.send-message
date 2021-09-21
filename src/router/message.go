package router

import (
	"github.com/Leonardo-Antonio/api.send-message/src/handler"
	"github.com/labstack/echo/v4"
)

func Message(app *echo.Echo) {
	handlerMessage := &handler.Message{}
	group := app.Group("/api/v1")
	group.POST("/users/message", handlerMessage.Send)
}
