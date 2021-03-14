package handler

import (
	"errors"
	"log"

	"github.com/profclems/go-dotenv"
)

type Config struct {
	Token string `yaml:"token"`
}

func SetToken(token string) {
	dotenv.SetConfigFile(".env")
	dotenv.LoadConfig()

	dotenv.Set("TOKEN", token)

	err := dotenv.Save()

	if err != nil {
		log.Fatal(err)
	}
}

func GetToken() (string, error) {
	err := dotenv.LoadConfig()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	token := dotenv.GetString("TOKEN")
	if token == "" {
		return "", errors.New("Token is empty")
	}
	return token, nil
}
