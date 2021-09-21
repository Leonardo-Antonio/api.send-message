package handler

import (
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.send-message/src/entity"
	"github.com/Leonardo-Antonio/api.send-message/src/helper/response"
	"github.com/Leonardo-Antonio/api.send-message/src/send"
	"github.com/labstack/echo/v4"
)

type Message struct{}

func (m *Message) Send(ctx echo.Context) error {
	var data entity.Message
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.New(response.ERR, "la estructura ingresada no es valida"))
	}

	if len(data.Name) == 0 || len(data.Email) == 0 || len(data.Text) == 0 {
		return ctx.JSON(http.StatusBadRequest, response.New(response.ERR, "los campos son obligatorios"))
	}

	if !strings.ContainsAny(data.Email, "@") {
		return ctx.JSON(http.StatusBadRequest, response.New(response.ERR, "ingrese un email valido"))
	}

	if err := send.Many(send.ReadTemplate("message", data), data.Email); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.New(response.ERR, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.New(response.MSJ, "se envio correctamente el mensaje"))
}
