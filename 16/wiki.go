package main

import (
	"fmt"
	"io/ioutil"

	"github.com/anonrootkit/16/utils"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := utils.GetFilename(page.Title)
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := utils.GetFilename(title)
	body, error := ioutil.ReadFile(filename)

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	homepage := Page{Title: "Home", Body: []byte("Hey there, welcome to this test page.")}
	homepage.save()

	homepageCopy, _ := loadPage(homepage.Title)
	fmt.Println(string(homepageCopy.Body))
}
