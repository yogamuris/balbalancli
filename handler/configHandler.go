package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/fatih/color"
	"github.com/profclems/go-dotenv"
)

func SetToken(token string) error {
	homedir := getHomedir()
	if _, err := os.Stat(homedir + "\\.balbalancli.env"); os.IsNotExist(err) {
		err := ioutil.WriteFile(homedir+"\\.balbalancli.env", []byte("TOKEN="), 0755)
		if err != nil {
			return err
		}
	}

	dotenv.SetConfigFile(homedir + "\\.balbalancli.env")
	dotenv.LoadConfig()

	dotenv.Set("TOKEN", token)

	err := dotenv.Save()

	if err != nil {
		return err
	}

	color.Green("Token added successfully!")
	return nil
}

func GetToken() (string, error) {
	homedir := getHomedir()
	dotenv.SetConfigFile(homedir + "\\.balbalancli.env")
	err := dotenv.LoadConfig()

	if err != nil {
		color.Red("Token is empty, set up your token with \"balbalancli config --token YOUR_TOKEN_VALUE\" command")
		return "", errors.New("Token is empty.")
	}

	token := dotenv.GetString("TOKEN")
	if token == "" {
		color.Red("Token is empty, set up your token with \"balbalancli config --token YOUR_TOKEN_VALUE\" command")
		return "", errors.New("Token is empty")
	}
	return token, nil
}

func getHomedir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	return homeDirectory
}
