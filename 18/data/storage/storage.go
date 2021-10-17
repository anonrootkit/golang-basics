package data

import (
	"io/ioutil"

	"github.com/anonrootkit/18/utils"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) Save() error {
	filename := utils.GetFileName(page.Title)
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func CreateTempPageAndSave(title string) {
	page := Page{Title: title, Body: []byte("No information provided for page : " + title)}
	page.Save()
}

func LoadPage(title string) (*Page, error) {
	filename := utils.GetFileName(title)
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
