package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yogamuris/balbalancli/model"
)

func GetLatestScore(league string) error {
	token, err := GetToken()
	if err != nil {
		return err
	}

	code := GetCompetitionCode(league)
	todayDate, previousDate := getTimeRange()

	url := fmt.Sprintf("https://api.football-data.org/v2/competitions/%d/matches?dateFrom=%s&dateTo=%s&status=FINISHED", code, previousDate, todayDate)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return errors.New("Error happened. Check your internet connection.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return GetResponseError(body)
	}

	var response model.ScoreResponse
	err = json.Unmarshal([]byte(body), &response)

	if err != nil {
		return errors.New("Error happened.")
	}

	printScore(&response)

	return nil
}

func getTimeRange() (string, string) {
	currentTime := time.Now()
	todayDate := currentTime.Format("2006-01-02")
	previousDate := currentTime.AddDate(0, 0, -1).Format("2006-01-02")

	return todayDate, previousDate
}
