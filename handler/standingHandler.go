package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yogamuris/balbalancli/model"
)

func GetStanding(league string) error {
	var err error
	if league == "All" {
		err = getAllStanding()
	} else {
		err = getStandingLeague(league)
	}

	if err != nil {
		return err
	}

	return nil
}

func getAllStanding() error {
	var err error
	for league := range GetAllCompetitionCode() {
		err = getStandingLeague(league)
		if err != nil {
			return err
		}
	}

	return nil
}

func getStandingLeague(league string) error {
	token, err := GetToken()

	if err != nil {
		return err
	}

	code := GetCompetitionCode(league)

	url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/standings?standingType=TOTAL", code)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response model.StandingResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return err
	}

	printStanding(&response)
	return nil
}
