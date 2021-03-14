package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func printStanding(response *StandingResponse) {
	fmt.Printf("Competition ==> %d %s %s\n", response.Competition.ID, response.Competition.Code, response.Competition.Name)
	for _, standing := range response.Standing {
		fmt.Printf("Standing ==> %s %s\n", standing.Stage, standing.Type)
		tables := standing.Tables

		standingTable := tablewriter.NewWriter(os.Stdout)
		standingTable.SetHeader([]string{"Position", "Team", "PG", "Form", "W", "D", "L", "P", "GF", "GA", "GD"})
		standingTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		standingTable.SetCenterSeparator("|")

		for _, table := range tables {
			standingTable.Append([]string{strconv.Itoa(table.Position), table.Team.Name, strconv.Itoa(table.PlayedGames), table.Form, strconv.Itoa(table.Won), strconv.Itoa(table.Draw), strconv.Itoa(table.Lost), strconv.Itoa(table.Points), strconv.Itoa(table.GoalsFor), strconv.Itoa(table.GoalsAgainst), strconv.Itoa(table.GoalDifference)})
		}

		standingTable.Render()
		fmt.Println()
	}
}

func printScore(response *ScoreResponse) {
	fmt.Printf("### %s ###\n", response.Competition.Name)

	scoreTable := tablewriter.NewWriter(os.Stdout)

	if len(response.Matches) == 0 {
		fmt.Println("No Match")
		fmt.Println()

	} else {
		for _, match := range response.Matches {
			scoreTable.Append([]string{match.Date.Format("02-01-2006"), match.HomeTeam.Name, strconv.Itoa(match.Scores.FullTime.HomeTeam), strconv.Itoa(match.Scores.FullTime.AwayTeam), match.AwayTeam.Name})
		}

		scoreTable.Render()
		fmt.Println()
	}
}
