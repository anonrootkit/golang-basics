package storage

import (
	"io/ioutil"

	"github.com/anonrootkit/19/data/network/model"
)

func SaveInFile(weather model.Weather, filename string) {
	ioutil.WriteFile(filename, []byte(weather.GetFormattedWeather()), 0666)
}
