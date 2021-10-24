package model

import (
	"fmt"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		Celcius    float32 `json:"temp_c"`
		Fahrenheit float32 `json:"temp_f"`
	} `json:"current"`
}

func (weather Weather) GetFormattedWeather() string {
	var weatherString = ""

	for i := 0; i < 25; i++ {
		weatherString += "*"
	}

	weatherString += "\n"

	weatherString +=
		"Region : " + weather.Location.Region + "\n" +
			"Country : " + weather.Location.Country + "\n" +
			"Temp : " + fmt.Sprint(weather.Current.Celcius) + " C" + "\n"

	for i := 0; i < 25; i++ {
		weatherString += "*"
	}

	return weatherString
}
