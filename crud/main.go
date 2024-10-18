package main

import (
	"crud/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/series", rest.GetAllSeries).Methods("GET")
	router.HandleFunc("/series/{id}", rest.GetSeriesById).Methods("GET")
	router.HandleFunc("/series", rest.CreateSeries).Methods("POST")
	router.HandleFunc("/series/{id}", rest.UpdateSeriesById).Methods("PUT")
	router.HandleFunc("/series/{id}", rest.DeleteSeriesById).Methods("DELETE")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal("Error in starting the server", err)
	}
}
