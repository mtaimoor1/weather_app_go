package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type apiConfigData struct {
	ApiKey string `json:"apiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		temp float64 `json:"temperature"`
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
func main() {
	cred, err := loadApiConfig(".apiConfig")
	if err != nil {
		log.Fatal("Error loading config")
	}
	fmt.Println(cred.ApiKey)
}
