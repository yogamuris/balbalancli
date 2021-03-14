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

type ScoreResponse struct {
	Competition model.Competition `json:"competition"`
	Matches     []model.Match     `json:"matches"`
}

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

	// fmt.Println(string(body))

	// body := `{"count":8,"filters":{},"competition":{"id":2021,"area":{"id":2072,"name":"England"},"name":"Premier League","code":"PL","plan":"TIER_ONE","lastUpdated":"2021-03-14T04:45:35Z"},"matches":[{"id":304034,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-13T12:30:00Z","status":"FINISHED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-14T00:00:02Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"DRAW","duration":"REGULAR","fullTime":{"homeTeam":0,"awayTeam":0},"halfTime":{"homeTeam":0,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":341,"name":"Leeds United FC"},"awayTeam":{"id":61,"name":"Chelsea FC"},"referees":[{"id":11487,"name":"Kevin Friend","role":"MAIN_REFEREE","nationality":"England"},{"id":11544,"name":"Simon Beck","role":"ASSISTANT_N1","nationality":"England"},{"id":11595,"name":"Adrian Holmes","role":"ASSISTANT_N2","nationality":"England"},{"id":11580,"name":"Anthony Taylor","role":"FOURTH_OFFICIAL","nationality":"England"},{"id":11494,"name":"Stuart Attwell","role":"VIDEO_ASSISTANT_REFEREE","nationality":"England"}]},{"id":304036,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-13T15:00:00Z","status":"FINISHED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-13T23:59:37Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"HOME_TEAM","duration":"REGULAR","fullTime":{"homeTeam":1,"awayTeam":0},"halfTime":{"homeTeam":1,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":354,"name":"Crystal Palace FC"},"awayTeam":{"id":74,"name":"West Bromwich Albion FC"},"referees":[{"id":11430,"name":"Simon Hooper","role":"MAIN_REFEREE","nationality":"England"},{"id":11521,"name":"Mark Scholes","role":"ASSISTANT_N1","nationality":"England"},{"id":11505,"name":"Derek Eaton","role":"ASSISTANT_N2","nationality":"England"},{"id":11503,"name":"Graham Scott","role":"FOURTH_OFFICIAL","nationality":"England"},{"id":11575,"name":"Mike Dean","role":"VIDEO_ASSISTANT_REFEREE","nationality":"England"}]},{"id":304031,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-13T17:30:00Z","status":"FINISHED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-14T03:05:59Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":1,"awayTeam":2},"halfTime":{"homeTeam":1,"awayTeam":2},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":62,"name":"Everton FC"},"awayTeam":{"id":328,"name":"Burnley FC"},"referees":[{"id":11567,"name":"Jonathan Moss","role":"MAIN_REFEREE","nationality":"England"},{"id":11531,"name":"Marc Perry","role":"ASSISTANT_N1","nationality":"England"},{"id":11431,"name":"Daniel Robathan","role":"ASSISTANT_N2","nationality":"England"},{"id":11520,"name":"Paul Tierney","role":"FOURTH_OFFICIAL","nationality":"England"},{"id":11551,"name":"Martin Atkinson","role":"VIDEO_ASSISTANT_REFEREE","nationality":"England"}]},{"id":304028,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-13T20:00:00Z","status":"FINISHED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-14T04:45:35Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":"AWAY_TEAM","duration":"REGULAR","fullTime":{"homeTeam":0,"awayTeam":3},"halfTime":{"homeTeam":0,"awayTeam":0},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":63,"name":"Fulham FC"},"awayTeam":{"id":65,"name":"Manchester City FC"},"referees":[{"id":11610,"name":"Andre Marriner","role":"MAIN_REFEREE","nationality":"England"},{"id":11611,"name":"Scott Ledger","role":"ASSISTANT_N1","nationality":"England"},{"id":11504,"name":"Simon Long","role":"ASSISTANT_N2","nationality":"England"},{"id":11446,"name":"Robert Jones","role":"FOURTH_OFFICIAL","nationality":"England"},{"id":11585,"name":"Craig Pawson","role":"VIDEO_ASSISTANT_REFEREE","nationality":"England"}]},{"id":304035,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-14T12:00:00Z","status":"SCHEDULED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-13T00:00:16Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":null,"duration":"REGULAR","fullTime":{"homeTeam":null,"awayTeam":null},"halfTime":{"homeTeam":null,"awayTeam":null},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":340,"name":"Southampton FC"},"awayTeam":{"id":397,"name":"Brighton & Hove Albion FC"},"referees":[]},{"id":304033,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-14T14:00:00Z","status":"SCHEDULED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-13T00:00:16Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":null,"duration":"REGULAR","fullTime":{"homeTeam":null,"awayTeam":null},"halfTime":{"homeTeam":null,"awayTeam":null},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":338,"name":"Leicester City FC"},"awayTeam":{"id":356,"name":"Sheffield United FC"},"referees":[]},{"id":304030,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-14T16:30:00Z","status":"SCHEDULED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-13T00:00:16Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":null,"duration":"REGULAR","fullTime":{"homeTeam":null,"awayTeam":null},"halfTime":{"homeTeam":null,"awayTeam":null},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":57,"name":"Arsenal FC"},"awayTeam":{"id":73,"name":"Tottenham Hotspur FC"},"referees":[]},{"id":304032,"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33},"utcDate":"2021-03-14T19:15:00Z","status":"SCHEDULED","matchday":28,"stage":"REGULAR_SEASON","group":"Regular Season","lastUpdated":"2021-03-13T00:00:16Z","odds":{"msg":"Activate Odds-Package in User-Panel to retrieve odds."},"score":{"winner":null,"duration":"REGULAR","fullTime":{"homeTeam":null,"awayTeam":null},"halfTime":{"homeTeam":null,"awayTeam":null},"extraTime":{"homeTeam":null,"awayTeam":null},"penalties":{"homeTeam":null,"awayTeam":null}},"homeTeam":{"id":66,"name":"Manchester United FC"},"awayTeam":{"id":563,"name":"West Ham United FC"},"referees":[]}]}`

	var response ScoreResponse
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
