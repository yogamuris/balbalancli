package model

// StandingResponse Model
type StandingResponse struct {
	Filter      string      `json:"filter"`
	Competition Competition `json:"competition"`
	Standing    []Standing  `json:"standings"`
}

type ScoreResponse struct {
	Competition Competition `json:"competition"`
	Matches     []Match     `json:"matches"`
}
