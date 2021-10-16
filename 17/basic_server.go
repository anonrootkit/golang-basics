package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hi there, I love Puca!")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
