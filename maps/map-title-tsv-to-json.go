package maps

import (
	"strings"
)

type Title struct {
	Tconst         string   `json: "tconst"`
	TitleType      string   `json: "titleType"`
	PrimaryTitle   string   `json: "primaryTitle"`
	OriginalTitle  string   `json: "originalTitle"`
	IsAdult        string   `json: "isAdult"`
	StartYear      string   `json: "startYear"`
	EndYear        string   `json: "endYear"`
	RuntimeMinutes string   `json: "runtimeMinutes"`
	Genres         []string `json: "genres"`
}

func CreateTitles(lines [][]string) []Title {

	// Loop through lines & turn into object
	var titles []Title
	for _, line := range lines {
		data := Title{
			Tconst:         line[0],
			TitleType:      line[1],
			PrimaryTitle:   line[2],
			OriginalTitle:  line[3],
			IsAdult:        line[4],
			StartYear:      line[5],
			EndYear:        line[6],
			RuntimeMinutes: line[7],
			Genres:         strings.Split(line[8], ","),
		}

		titles = append(titles, data)
	}
	return titles
}
