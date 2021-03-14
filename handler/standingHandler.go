package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yogamuris/balbalancli/model"
)

func GetStanding(league string) {
	if league == "All" {
		getAllStanding()
	} else {
		getStandingLeague(league)
	}
}

func getAllStanding() {
	for league := range GetAllCompetitionCode() {
		getStandingLeague(league)
	}
}

func getStandingLeague(league string) {
	token, err := GetToken()

	if err != nil {
		fmt.Println(err)
	}

	code := GetCompetitionCode(league)

	url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/standings?standingType=TOTAL", code)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response model.StandingResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println(err)
	}

	printStanding(&response)
}
