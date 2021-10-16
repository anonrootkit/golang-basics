package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	for id := range ids {
		fmt.Printf("%d %d\n", id, id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	mapUse()
}

func mapUse() {
	names := map[string]string{"Ankit": "Boljub", "Puca": "Pucaa"}

	for name, value := range names {
		fmt.Print(name + " " + value + "\n")
	}
}
