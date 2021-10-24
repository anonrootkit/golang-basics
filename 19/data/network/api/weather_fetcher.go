package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anonrootkit/19/data/network/model"
)

const base_url = "https://api.weatherapi.com/v1/"
const key = "91e974c06e0d40caadd115413211910"

func FetchCurrentWeather(state string) model.Weather {
	currentWeatherUrl := base_url + "current.json?key=" + key + "&q=" + state

	fmt.Println(currentWeatherUrl)

	response, nwError := http.Get(currentWeatherUrl)

	if nwError != nil {
		log.Fatalln("Error fetching data from network \n\n Error : " + nwError.Error())
	}

	defer response.Body.Close()

	var weather model.Weather
	parsingErr := json.NewDecoder(response.Body).Decode(&weather)

	if parsingErr != nil {
		log.Fatalln("Error occurred while parsing response \n\n Error : " + parsingErr.Error())
	}

	return weather
}
