package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var key = "5948fd9227ade17558f52fd9554e0cb9"
var town = ""

func main() {
	getCity()
}

func getCity() {
	fmt.Println("Введите название города: ")
	fmt.Scan(&town)
	if len(town) == 0 {
		getCity()
	} else {
		getWeather()
		getCity()
	}
}

func getWeather() {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + town + "&appid=" + key)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	var weatherResp WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Получение данных о погоде
	weather := weatherResp.Weather[0]
	mainInfo := weatherResp.Main

	fmt.Println("Main weather:", weather.Main)
	fmt.Println("Description:", weather.Description)
	fmt.Println("Temperature (K):", mainInfo.Temperature)
	fmt.Println("Pressure:", mainInfo.Pressure)
	fmt.Println("Humidity:", mainInfo.Humidity)

	getCity()
}
