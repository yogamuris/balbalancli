package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/yogamuris/balbalancli/model"
)

func printStanding(response *model.StandingResponse) {
	fmt.Printf("### %s ###\n\n", response.Competition.Name)

	if len(response.Standing) == 0 {
		fmt.Println("No Football Standing")
	} else {
		for _, standing := range response.Standing {
			tables := standing.Tables

			standingTable := tablewriter.NewWriter(os.Stdout)
			standingTable.SetHeader([]string{"Position", "Team", "PG", "Form", "W", "D", "L", "P", "GF", "GA", "GD"})

			for _, table := range tables {
				standingTable.Append([]string{strconv.Itoa(table.Position), table.Team.Name, strconv.Itoa(table.PlayedGames), table.Form, strconv.Itoa(table.Won), strconv.Itoa(table.Draw), strconv.Itoa(table.Lost), strconv.Itoa(table.Points), strconv.Itoa(table.GoalsFor), strconv.Itoa(table.GoalsAgainst), strconv.Itoa(table.GoalDifference)})
			}

			standingTable.Render()
			fmt.Println()
		}
	}
}

func printScore(response *model.ScoreResponse) {
	fmt.Printf("### %s ###\n", response.Competition.Name)

	scoreTable := tablewriter.NewWriter(os.Stdout)

	if len(response.Matches) == 0 {
		fmt.Println("No Football Match")
		fmt.Println()

	} else {
		for _, match := range response.Matches {
			scoreTable.Append([]string{match.Date.Format("02-01-2006"), match.HomeTeam.Name, strconv.Itoa(match.Scores.FullTime.HomeTeam), strconv.Itoa(match.Scores.FullTime.AwayTeam), match.AwayTeam.Name})
		}
		scoreTable.Render()
		fmt.Println()
	}
}

func PrintLeagueList() {
	leagueTable := tablewriter.NewWriter(os.Stdout)
	leagueTable.SetHeader([]string{"Code", "League"})
	leagueTable.Append([]string{"PL", "English Premier League"})
	leagueTable.Append([]string{"PD", "Primera Division La Liga"})
	leagueTable.Append([]string{"SA", "Serie A"})
	leagueTable.Append([]string{"BL1", "Bundesliga"})
	leagueTable.Append([]string{"FL1", "Ligue 1"})
	leagueTable.Append([]string{"DED", "Eredivisie"})
	leagueTable.Append([]string{"PPL", "Primeira Liga"})
	leagueTable.Append([]string{"CL", "UEFA Champions League"})

	leagueTable.Render()
}
