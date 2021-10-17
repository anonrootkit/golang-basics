package utils

import (
	"errors"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func GetPageTitle(writer http.ResponseWriter, request *http.Request) (string, error) {
	matches := validPath.FindStringSubmatch(request.URL.Path)

	if matches == nil {
		http.NotFound(writer, request)
		return "", errors.New("invalid url")
	}

	return matches[2], nil
}

func GetFileName(title string) string {
	return title + ".txt"
}
