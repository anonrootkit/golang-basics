package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, requester *http.Request) {
	fmt.Fprintf(writer, "<h1>Hello world</h1>")
}

func about(writer http.ResponseWriter, requester *http.Request) {
	fmt.Fprintf(writer, "My name is Ankit")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", nil)
}
