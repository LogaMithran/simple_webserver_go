package rest

import (
	"crud/types"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var constDate, _ = time.Parse("Jan 01, 1999", "01-02-2001")
var constDob, _ = time.Parse("Jan 01, 1999", "01-01-2001")

var series = []types.Series{
	{
		Id:           1,
		Series_name:  "The Walking Dead",
		Release_date: constDate,
		Director: &types.Director{
			Director_name: "Mithran",
			Date_of_birth: constDob,
		},
		Season: []types.Season{{Season_number: 1, Number_of_episodes: 10}, {Season_number: 2, Number_of_episodes: 13}, {Season_number: 3, Number_of_episodes: 15}},
	},
	{
		Id:           2,
		Series_name:  "Friends",
		Release_date: constDate,
		Director: &types.Director{
			Director_name: "Rishi",
			Date_of_birth: constDob,
		},
		Season: []types.Season{{Season_number: 1, Number_of_episodes: 22}, {Season_number: 2, Number_of_episodes: 23}, {Season_number: 3, Number_of_episodes: 20}},
	},
	{
		Id:           3,
		Series_name:  "The office",
		Release_date: constDate,
		Director: &types.Director{
			Director_name: "Rob",
			Date_of_birth: constDob,
		},
		Season: []types.Season{{Season_number: 1, Number_of_episodes: 10}, {Season_number: 2, Number_of_episodes: 13}, {Season_number: 3, Number_of_episodes: 15}},
	},
}

func GetAllSeries(response http.ResponseWriter, request *http.Request) {
	log.Print("This is get all the movies")
	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response).Encode(series)
}

func GetSeriesById(response http.ResponseWriter, request *http.Request) {
	log.Print("This is get movies by id")
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	target, _ := strconv.Atoi(params["id"])
	for _, serie := range series {
		if serie.Id == target {
			json.NewEncoder(response).Encode(serie)
			return
		}
	}
}

func CreateSeries(response http.ResponseWriter, request *http.Request) {
	log.Print("This is create movies")
	response.Header().Set("content-type", "application/json")
	var serie types.Series

	_ = json.NewDecoder(request.Body).Decode(&serie)
	serie.Id = rand.Int()
	series = append(series, serie)
	json.NewEncoder(response).Encode(series)
}

func UpdateSeriesById(response http.ResponseWriter, request *http.Request) {
	log.Print("This is update movies by id")
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	target, _ := strconv.Atoi(params["id"])

	for i, serie := range series {
		if serie.Id == target {
			series = append(series[:i], series[i+1:]...)
			serie.Id = target
			json.NewDecoder(request.Body).Decode(&serie)
			series = append(series, serie)
			json.NewEncoder(response).Encode(series)
			return
		}
	}
}

func DeleteSeriesById(response http.ResponseWriter, request *http.Request) {
	log.Print("This is delete movies by id")
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	target, _ := strconv.Atoi(params["id"])

	for i, serie := range series {
		if serie.Id == target {
			series = append(series[:i], series[i+1:]...)
			json.NewEncoder(response).Encode(series)
			return
		}
	}
}
