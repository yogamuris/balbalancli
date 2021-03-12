package handler

var competitionCode = map[string]int{
	"PL":  2021,
	"SA":  2019,
	"CL":  2001,
	"FL1": 2015,
	"BL1": 2002,
	"DED": 2003,
	"PPL": 2017,
	"PD":  2014,
}

// GetCompetitionCode return the competition ID
func GetCompetitionCode(league string) int {
	return competitionCode[league]
}
