package weatherapi

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/anonrootkit/20/data/network"
)

func FetchWeather(state string, days int) []byte {
	forecastPath := "forecast.json" + "?q=" + state + "&days=" + strconv.Itoa(days)

	response, err := network.FetchData(forecastPath)

	if err != nil {
		log.Fatalf("Error occurred while fetcing weather for %s and for %d days\nError : %s", state, days, err.Error())
	}

	defer response.Body.Close()

	var data = map[string]interface{}{}

	parsingErr := json.NewDecoder(response.Body).Decode(&data)

	if parsingErr != nil {
		log.Fatalln("Error occurred while parsing response \n\n Error : " + parsingErr.Error())
	}

	jsonData, _ := json.Marshal(data)

	return jsonData
}
