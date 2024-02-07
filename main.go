package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	ApiKey string `json:"apiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Unable to read the config file")
	}
	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal("Unable to Unmarshal the bytes")
	}
	return c, nil
}
func query_api(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.ApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}
func main() {
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query_api(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8080", nil)
}
