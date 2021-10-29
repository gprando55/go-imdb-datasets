package main

import (
	"encoding/json"
	"fmt"
	"log"

	Maps "api/maps"
	Services "api/services"
)

const ratingFileGz = "./rating.gz"
const titlesFileGz = "./titles.gz"

const ratingUrl = "https://datasets.imdbws.com/title.ratings.tsv.gz"
const titlesUrl = "https://datasets.imdbws.com/title.basics.tsv.gz"

const ratingFileTsv = "./rating.tsv"

const titlesFileTsv = "./titles.tsv"

const ratingFileJson = "./rating.json"
const titlesFileJson = "./titles.json"

func main() {
	//  download files
	Services.DownloadFile(ratingFileGz, ratingUrl)
	Services.DownloadFile(titlesFileGz, titlesUrl)

	//  unzip files
	Services.UnGzip(ratingFileGz, ratingFileTsv)
	Services.UnGzip(titlesFileGz, titlesFileTsv)

	fmt.Println("Download finished")

	// open file
	dataTitles, errReadTitles := Services.ReadTsv(titlesFileTsv)
	dataRatings, errReadRating := Services.ReadTsv(ratingFileTsv)
	if errReadRating != nil || errReadTitles != nil {
		log.Fatal(errReadTitles, errReadRating)
	}

	fmt.Println("read files finished")

	titles := Maps.CreateTitles(dataTitles)
	ratings := Maps.CreateRatings(dataRatings)

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonDataTitles, errJsonTitles := json.MarshalIndent(titles, "", "  ")
	if errJsonTitles != nil {
		log.Fatal(errJsonTitles)
	}

	fmt.Println(string(jsonDataTitles))

	Services.SaveJsonFile(titlesFileJson, string(jsonDataTitles))

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonDataRatings, errJsonRatings := json.MarshalIndent(ratings, "", "  ")
	if errJsonRatings != nil {
		log.Fatal(errJsonRatings)
	}

	fmt.Println(string(jsonDataRatings))

	Services.SaveJsonFile(ratingFileJson, string(jsonDataRatings))

	Services.RemoveFile(ratingFileGz)
	Services.RemoveFile(titlesFileGz)
	Services.RemoveFile(ratingFileTsv)
	Services.RemoveFile(titlesFileTsv)
}
