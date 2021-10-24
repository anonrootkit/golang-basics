package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anonrootkit/19/data/network/api"
	"github.com/anonrootkit/19/data/storage"
	"github.com/anonrootkit/19/data/storage/utils"
)

var inputSize = 1000000

func main() {
	fmt.Print("Enter the state name : ")
	var state string

	input := bufio.NewScanner(os.Stdin)
	input.Buffer(make([]byte, inputSize), inputSize)

	input.Scan()

	fmt.Sscan(input.Text(), &state)

	weather := api.FetchCurrentWeather(state)

	storage.SaveInFile(weather, utils.GetWeatherFile())

	fmt.Println("Data saved in " + utils.GetWeatherFile())
}
