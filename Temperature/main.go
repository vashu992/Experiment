package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)
type Current struct {
	Temperature float64  `json:"temperature"`
}
type Response struct {
	Current Current
}

func  main() {
	r := mux.NewRouter()
	r.HandleFunc("/temperature", getTemperature).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000,",r))
}

func getTemperature(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("city")
	url := fmt.Sprintf("https://api.weatherstack.com/current?access_key=6b8adcd7a94db316fb0daf0bc7e759&query=%s",&query)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var wresp Response
	err = json.NewDecoder(response.Body).Decode(&wresp)
	if err != nil {
		log.Fatal(err) 
	}
	w.Header().Set("Contect-Type","application/json")
	json.NewEncoder(w).Encode(wresp.Current.Temperature)
} 
