package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/yogamuris/balbalancli/model"
)

func GetLatestScore(league string) {
	if league == "All" {
		getAllLeagueLatestScore()
	} else {
		getLeagueLatestScore(league)
	}
}

func getAllLeagueLatestScore() {
	for league := range GetAllCompetitionCode() {
		getLeagueLatestScore(league)
	}
}

func getLeagueLatestScore(league string) {
	token, err := GetToken()
	if err != nil {
		log.Fatal(err)
	}

	code := GetCompetitionCode(league)
	todayDate, previousDate := getTimeRange()

	url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/matches?dateFrom=%s&dateTo=%s&status=FINISHED", code, previousDate, todayDate)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response model.ScoreResponse
	err = json.Unmarshal([]byte(body), &response)

	if err != nil {
		fmt.Println(err)
	}

	printScore(&response)
}

func getTimeRange() (string, string) {
	currentTime := time.Now()
	todayDate := currentTime.Format("2006-01-02")
	previousDate := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	return todayDate, previousDate
}
