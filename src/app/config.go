package app

import (
	"log"

	"github.com/Leonardo-Antonio/api.send-message/src/env"
	"github.com/joho/godotenv"
)

func Config() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	env.GetEnv()
}
