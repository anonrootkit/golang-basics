package storage

import (
	"io"
	"log"
	"os"
)

func StoreWeatherInStorage(filename string, weatherData string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer file.Close()

	writer := io.MultiWriter(os.Stdout, file)

	writer.Write([]byte(weatherData))

}
