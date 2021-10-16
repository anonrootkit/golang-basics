package main

import (
	"fmt"
)

func main() {
	namesAndEmails := make(map[string]string)
	namesAndPh := map[string]string{"Ankit": "ankitankit", "Boljub": "BoljubBoljub"}

	// assign kv

	namesAndEmails["ankit"] = "ankit.gup20@gmail.com"
	namesAndEmails["boljub"] = "boljub.lftbsj@gmail.com"

	fmt.Println(namesAndEmails)
	fmt.Println(len(namesAndEmails))

	delete(namesAndEmails, "boljub")

	fmt.Println(namesAndEmails)
	fmt.Println(len(namesAndEmails))
	fmt.Println(namesAndPh)
}
