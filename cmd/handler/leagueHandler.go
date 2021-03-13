package handler

var competitionCode = map[string]int{
	"PL":  2021,
	"PD":  2014,
	"SA":  2019,
	"BL1": 2002,
	"FL1": 2015,
	"DED": 2003,
	"PPL": 2017,
	"CL":  2001,
}

// GetCompetitionCode return the competition ID
func GetCompetitionCode(league string) int {
	return competitionCode[league]
}

func GetAllCompetitionCode() map[string]int {
	return competitionCode
}
