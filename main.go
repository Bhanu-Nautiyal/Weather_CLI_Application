package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("ERROR LOADING THE ENV FILE")
		return
	}

	// Getting the apilink and apikey from .env file
	apiLink := os.Getenv("CALLER_LINK")
	apiKey := os.Getenv("API_KEY")
	aqi := "yes"
	// Take the Location as User Input
	usr_args := os.Args[1:]
	location := ""

	if len(usr_args) == 0 {
		fmt.Println("Cant work without city name")
		return
	}

	if len(usr_args) == 1 {
		location = usr_args[0]
	} else {
		fmt.Println("You entered more than one cities, Please just give 1 city name")
		return
	}

	// Created the apiUrl
	apiUrl := apiLink + apiKey + "&q=" + location + "&" + aqi
	req, err := http.Get(apiUrl)

	// Here I just took body of 'req' away from it
	// body := req.Body
	// resp, err := io.ReadAll(body)

	// if err != nil {
	// 	fmt.Println("Cant Parse the body of response")
	// 	return
	// }

	// Take the JSON and save it like map
	var data map[string]interface{}

	body := req.Body
	resp, err := io.ReadAll(body)

	if err != nil {
		fmt.Println("Cant parse the body")
		return
	}

	_ = json.Unmarshal(resp, &data)

	// Now I want specific part of the JSON response to get displayed at user terminal

	fmt.Println(data)
}
