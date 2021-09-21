package env

import (
	"os"
	"sync"
)

type credentials struct {
	Email    string
	Password string
}

var Credentials *credentials
var once sync.Once
var Port string

func GetEnv() {
	once.Do(func() {
		Credentials = &credentials{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD_EMAIL"),
		}
		Port = os.Getenv("PORT")
	})
}
