package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/yogamuris/balbalancli/model"
)

// StandingResponse Model
type StandingResponse struct {
	Filter      string            `json:"filter"`
	Competition model.Competition `json:"competition"`
	Standing    []model.Standing  `json:"standings"`
}

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

	// body := `{"filters":{},"competition":{"id":2021,"area":{"id":2072,"name":"England"},"name":"Premier League","code":"PL","plan":"TIER_ONE","lastUpdated":"2021-03-11T02:20:01Z"},"season":{"id":619,"startDate":"2020-09-12","endDate":"2021-05-23","currentMatchday":33,"winner":null},"standings":[{"stage":"REGULAR_SEASON","type":"TOTAL","group":null,"table":[{"position":1,"team":{"id":65,"name":"Manchester City FC","crestUrl":"https://crests.football-data.org/65.svg"},"playedGames":29,"form":"W,L,W,W,W","won":21,"draw":5,"lost":3,"points":68,"goalsFor":61,"goalsAgainst":21,"goalDifference":40},{"position":2,"team":{"id":66,"name":"Manchester United FC","crestUrl":"https://crests.football-data.org/66.svg"},"playedGames":28,"form":"W,D,D,W,D","won":15,"draw":9,"lost":4,"points":54,"goalsFor":55,"goalsAgainst":32,"goalDifference":23},{"position":3,"team":{"id":338,"name":"Leicester City FC","crestUrl":"https://crests.football-data.org/338.svg"},"playedGames":28,"form":"W,D,L,W,W","won":16,"draw":5,"lost":7,"points":53,"goalsFor":48,"goalsAgainst":32,"goalDifference":16},{"position":4,"team":{"id":61,"name":"Chelsea FC","crestUrl":"https://crests.football-data.org/61.svg"},"playedGames":28,"form":"W,W,D,D,W","won":14,"draw":8,"lost":6,"points":50,"goalsFor":44,"goalsAgainst":25,"goalDifference":19},{"position":5,"team":{"id":563,"name":"West Ham United FC","crestUrl":"https://crests.football-data.org/563.svg"},"playedGames":27,"form":"W,L,W,W,D","won":14,"draw":6,"lost":7,"points":48,"goalsFor":42,"goalsAgainst":31,"goalDifference":11},{"position":6,"team":{"id":62,"name":"Everton FC","crestUrl":"https://crests.football-data.org/62.svg"},"playedGames":27,"form":"L,W,W,W,L","won":14,"draw":4,"lost":9,"points":46,"goalsFor":39,"goalsAgainst":35,"goalDifference":4},{"position":7,"team":{"id":73,"name":"Tottenham Hotspur FC","crestUrl":"https://crests.football-data.org/73.svg"},"playedGames":27,"form":"W,W,W,L,L","won":13,"draw":6,"lost":8,"points":45,"goalsFor":46,"goalsAgainst":28,"goalDifference":18},{"position":8,"team":{"id":64,"name":"Liverpool FC","crestUrl":"https://crests.football-data.org/64.svg"},"playedGames":28,"form":"L,L,W,L,L","won":12,"draw":7,"lost":9,"points":43,"goalsFor":47,"goalsAgainst":36,"goalDifference":11},{"position":9,"team":{"id":58,"name":"Aston Villa FC","crestUrl":"https://crests.football-data.org/58.svg"},"playedGames":26,"form":"D,L,W,L,D","won":12,"draw":4,"lost":10,"points":40,"goalsFor":38,"goalsAgainst":27,"goalDifference":11},{"position":10,"team":{"id":57,"name":"Arsenal FC","crestUrl":"https://crests.football-data.org/57.svg"},"playedGames":27,"form":"D,W,L,W,L","won":11,"draw":5,"lost":11,"points":38,"goalsFor":35,"goalsAgainst":28,"goalDifference":7},{"position":11,"team":{"id":341,"name":"Leeds United FC","crestUrl":"https://crests.football-data.org/341.svg"},"playedGames":27,"form":"L,L,W,L,L","won":11,"draw":2,"lost":14,"points":35,"goalsFor":43,"goalsAgainst":46,"goalDifference":-3},{"position":12,"team":{"id":76,"name":"Wolverhampton Wanderers FC","crestUrl":"https://crests.football-data.org/76.svg"},"playedGames":28,"form":"D,L,D,W,W","won":9,"draw":8,"lost":11,"points":35,"goalsFor":28,"goalsAgainst":37,"goalDifference":-9},{"position":13,"team":{"id":354,"name":"Crystal Palace FC","crestUrl":"https://crests.football-data.org/354.svg"},"playedGames":28,"form":"L,D,D,W,L","won":9,"draw":7,"lost":12,"points":34,"goalsFor":30,"goalsAgainst":47,"goalDifference":-17},{"position":14,"team":{"id":340,"name":"Southampton FC","crestUrl":"https://crests.football-data.org/340.svg"},"playedGames":28,"form":"L,W,L,L,D","won":9,"draw":6,"lost":13,"points":33,"goalsFor":35,"goalsAgainst":49,"goalDifference":-14},{"position":15,"team":{"id":328,"name":"Burnley FC","crestUrl":"https://crests.football-data.org/328.svg"},"playedGames":28,"form":"D,D,L,D,D","won":7,"draw":9,"lost":12,"points":30,"goalsFor":20,"goalsAgainst":36,"goalDifference":-16},{"position":16,"team":{"id":67,"name":"Newcastle United FC","crestUrl":"https://crests.football-data.org/67.svg"},"playedGames":27,"form":"D,D,L,L,W","won":7,"draw":6,"lost":14,"points":27,"goalsFor":27,"goalsAgainst":44,"goalDifference":-17},{"position":17,"team":{"id":397,"name":"Brighton & Hove Albion FC","crestUrl":"https://crests.football-data.org/397.svg"},"playedGames":27,"form":"L,L,L,D,D","won":5,"draw":11,"lost":11,"points":26,"goalsFor":27,"goalsAgainst":35,"goalDifference":-8},{"position":18,"team":{"id":63,"name":"Fulham FC","crestUrl":"https://crests.football-data.org/63.svg"},"playedGames":28,"form":"W,L,D,W,D","won":5,"draw":11,"lost":12,"points":26,"goalsFor":22,"goalsAgainst":33,"goalDifference":-11},{"position":19,"team":{"id":74,"name":"West Bromwich Albion FC","crestUrl":"https://crests.football-data.org/74.svg"},"playedGames":28,"form":"D,L,W,D,D","won":3,"draw":9,"lost":16,"points":18,"goalsFor":20,"goalsAgainst":56,"goalDifference":-36},{"position":20,"team":{"id":356,"name":"Sheffield United FC","crestUrl":"https://crests.football-data.org/356.svg"},"playedGames":28,"form":"L,W,L,L,L","won":4,"draw":2,"lost":22,"points":14,"goalsFor":16,"goalsAgainst":45,"goalDifference":-29}]}]}`

	var response StandingResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println(err)
	}

	printStanding(&response)
}

func printStanding(response *StandingResponse) {
	fmt.Printf("Competition ==> %d %s %s\n", response.Competition.ID, response.Competition.Code, response.Competition.Name)
	for _, standing := range response.Standing {
		fmt.Printf("Standing ==> %s %s\n", standing.Stage, standing.Type)
		tables := standing.Tables

		t := tablewriter.NewWriter(os.Stdout)
		t.SetHeader([]string{"Position", "Team", "PG", "Form", "W", "D", "L", "P", "GF", "GA", "GD"})
		t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		t.SetCenterSeparator("|")

		for _, table := range tables {
			t.Append([]string{strconv.Itoa(table.Position), table.Team.Name, strconv.Itoa(table.PlayedGames), table.Form, strconv.Itoa(table.Won), strconv.Itoa(table.Draw), strconv.Itoa(table.Lost), strconv.Itoa(table.Points), strconv.Itoa(table.GoalsFor), strconv.Itoa(table.GoalsAgainst), strconv.Itoa(table.GoalDifference)})
		}

		t.Render()

		fmt.Println()
	}
}
