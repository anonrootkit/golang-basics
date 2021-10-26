package handler

import (
	"fmt"
	"github.com/anonrootkit/20/data/storage"
	"github.com/anonrootkit/20/model"
	"html/template"
	"net/http"
	"strconv"

	weatherapi "github.com/anonrootkit/20/data/network/weather_api"
)

const port = ":8090"

func InitHandlers() {
	indexHandler()
	fetchHandler()

	listenAndServe()
}

func listenAndServe() {
	http.ListenAndServe(port, nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	indexTemplate, err := template.ParseFiles("20/html/index.html")

	if err != nil {
		http.Error(writer, err.Error(), request.Response.StatusCode)
	}

	err = indexTemplate.ExecuteTemplate(writer, "index.html", nil)

	if err != nil {
		http.Error(writer, err.Error(), request.Response.StatusCode)
	}

}

func indexHandler() {
	http.HandleFunc("/", index)
	fmt.Println("Index handler started")
}

func fetch(writer http.ResponseWriter, request *http.Request) {
	state := request.FormValue("state")
	days, _ := strconv.Atoi(request.FormValue("days"))

	data := weatherapi.FetchWeather(state, days)
	weatherTemplate := template.New("weatherTemplate").Funcs(indexIncrementerFuncMap())
	weatherTemplate, err := weatherTemplate.ParseFiles("20/html/weather.html")
	if err != nil {
		http.Error(writer, err.Error(), request.Response.StatusCode)
	}

	weather := model.GetWeather(data)

	err = weatherTemplate.ExecuteTemplate(writer, "weather.html", weather)
	if err != nil {
		http.Error(writer, err.Error(), request.Response.StatusCode)
	}

	storage.StoreWeatherInStorage(storage.GetFileName(weather.Location.LOCATION), weather.GetFormattedWeather())
}

func indexIncrementerFuncMap() template.FuncMap {
	return template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
}

func fetchHandler() {
	http.HandleFunc("/fetch", fetch)
	fmt.Println("Fetch handler started")
}
