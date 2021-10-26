package model

import (
	"encoding/json"
	"fmt"
	"time"
)

//goland:noinspection ALL
type Weather struct {
	Current struct {
		TEMP_CEL float32 `json:"temp_c"`
		TEMP_FAH float32 `json:"temp_f"`

		WIND_KPH float32 `json:"wind_kph"`
		WIND_MPH float32 `json:"wind_mph"`
	} `json:"current"`

	Forecast struct {
		Days []struct {
			DAY struct {
				TEMP_CEL float32 `json:"avgtemp_c"`
				TEMP_FAH float32 `json:"avgtemp_f"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`

	Location struct {
		COUNTRY  string `json:"country"`
		STATE    string `json:"region"`
		LOCATION string `json:"name"`
	} `json:"location"`
}

func (weather Weather) GetFormattedWeather() string {
	formattedWeather := ""

	for i := 1; i <= 25; i++ {
		formattedWeather += "*"
	}

	formattedWeather += "\n"
	formattedWeather += time.Now().Format(time.RFC1123)
	formattedWeather += "\n"

	formattedWeather += "Region : " + weather.Location.LOCATION + ", " + weather.Location.COUNTRY + "\n"
	formattedWeather += "Temp : " + fmt.Sprint(weather.Current.TEMP_CEL) + " C ~ " + fmt.Sprint(weather.Current.TEMP_FAH) + " F\n\n"
	formattedWeather += "Forecast for " + fmt.Sprint(len(weather.Forecast.Days)) + " days : \n"

	for idx, day := range weather.Forecast.Days {
		formattedWeather += "Day " + fmt.Sprint(idx+1) + " : " + fmt.Sprint(day.DAY.TEMP_CEL) + " C ~ " + fmt.Sprint(day.DAY.TEMP_FAH) + " F\n"
	}

	for i := 1; i <= 25; i++ {
		formattedWeather += "*"
	}

	return formattedWeather
}

func GetWeather(data []byte) Weather {
	var weather Weather
	err := json.Unmarshal(data, &weather)

	if err != nil {
		fmt.Print(err.Error())
	}

	return weather
}
