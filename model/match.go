package model

import "time"

type Match struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"utcDate"`
	Status   string    `json:"status"`
	Matchday int       `json:"matchday"`
	Stage    string    `json:"stage"`
	Group    string    `json:"group"`
	Scores   Score     `json:"score"`
	HomeTeam Team      `json:"homeTeam"`
	AwayTeam Team      `json:"awayTeam"`
}

type Score struct {
	Winner    string    `json:"winner"`
	Duration  string    `json:"duration"`
	FullTime  TeamScore `json:"fullTime"`
	HalfTime  TeamScore `json:"halfTime"`
	ExtraTime TeamScore `json:"extraTime"`
	Penalties TeamScore `json:"penalties"`
}

type TeamScore struct {
	HomeTeam int `json:"homeTeam"`
	AwayTeam int `json:"awayTeam"`
}
