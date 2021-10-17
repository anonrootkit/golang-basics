package main

import (
	"github.com/anonrootkit/18/data/network"
)

func main() {
	network.IndexHandler()
	network.ViewPageHandler()
	network.EditPageHandler()
	network.SavePageHandler()
	network.ListenAndServe()
}
