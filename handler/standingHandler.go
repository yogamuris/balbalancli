package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yogamuris/balbalancli/model"
)

func GetStanding(league string) error {
	token, err := GetToken()

	if err != nil {
		return errors.New("Error happened.")
	}

	code := GetCompetitionCode(league)

	url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/standings?standingType=TOTAL", code)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return errors.New("Error happened. Check your internet connection.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response model.StandingResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return errors.New("Error happened.")
	}

	printStanding(&response)
	return nil
}
