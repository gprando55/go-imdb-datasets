package maps

type Rating struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Tconst        string `json:"tconst"`
	AverageRating string `json:"averageRating"`
	NumVotes      string `json:"numVotes"`
}

func CreateRatings(lines [][]string) []Rating {

	// Loop through lines & turn into object
	var ratings []Rating
	for _, line := range lines {
		data := Rating{
			Tconst:        line[0],
			AverageRating: line[1],
			NumVotes:      line[2],
		}

		ratings = append(ratings, data)
	}
	return ratings
}
