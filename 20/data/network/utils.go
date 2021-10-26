package network

import (
	"net/http"
)

const base_url = "https://api.weatherapi.com/v1/"
const key = "&key=91e974c06e0d40caadd115413211910"

func getUrlWithKey(path string) string {
	return base_url + path + key
}

func FetchData(path string) (*http.Response, error) {
	return http.Get(getUrlWithKey(path))
}
