package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/profclems/go-dotenv"
)

func SetToken(token string) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		err := ioutil.WriteFile(".env", []byte("TOKEN="), 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	dotenv.SetConfigFile(".env")
	dotenv.LoadConfig()

	dotenv.Set("TOKEN", token)

	err := dotenv.Save()

	if err != nil {
		log.Fatal(err)
	}

	color.Green("Token added successfully!")
}

func GetToken() (string, error) {
	err := dotenv.LoadConfig()

	if err != nil {
		color.Red("Token is empty, set up your token with \"balbalancli config --token YOUR_TOKEN_VALUE\" command")
		return "", errors.New("Token is empty.")
	}

	token := dotenv.GetString("TOKEN")
	if token == "" {
		return "", errors.New("Token is empty")
	}
	return token, nil
}
