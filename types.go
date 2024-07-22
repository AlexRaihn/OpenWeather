package main

type WeatherResponse struct {
	Weather []WeatherInfo `json:"weather"`
	Main    MainInfo      `json:"main"`
}

type WeatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type MainInfo struct {
	Temperature float64 `json:"temp"`
	Pressure    int     `json:"pressure"`
	Humidity    int     `json:"humidity"`
}
