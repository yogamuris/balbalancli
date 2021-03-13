package model

// Standing model
type Standing struct {
	Stage  string  `json:"stage"`
	Type   string  `json:"type"`
	Tables []Table `json:"table"`
}

// Team model
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Table model
type Table struct {
	Position       int    `json:"position"`
	Team           Team   `json:"team"`
	PlayedGames    int    `json:"playedGames"`
	Form           string `json:"form"`
	Won            int    `json:"won"`
	Draw           int    `json:"draw"`
	Lost           int    `json:"lost"`
	Points         int    `json:"points"`
	GoalsFor       int    `json:"goalsFor"`
	GoalsAgainst   int    `json:"goalsAgainst"`
	GoalDifference int    `json:"goalDifference"`
}
