package main

import (
	"github.com/Leonardo-Antonio/api.send-message/src/app"
)

func main() {
	app.Config()
	app := app.NewApp()
	app.Middlewares()
	app.Routers()
	app.Listeing()
}
