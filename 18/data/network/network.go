package network

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	data "github.com/anonrootkit/18/data/storage"
	"github.com/anonrootkit/18/utils"
)

var port = ":8080"
var templates = template.Must(template.ParseFiles("html/view.html", "html/edit.html"))

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "This is series : %d", 18)
}

func IndexHandler() {
	http.HandleFunc("/", index)
	fmt.Println("Starting index handler")
}

func viewPage(writer http.ResponseWriter, request *http.Request) {
	title, _ := utils.GetPageTitle(writer, request)
	page, err := data.LoadPage(title)

	if err != nil {
		data.CreateTempPageAndSave(title)
		viewPage(writer, request)
	} else {
		err := templates.ExecuteTemplate(writer, "view.html", page)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ViewPageHandler() {
	http.HandleFunc("/view/", viewPage)
	fmt.Println("Starting view page handler")
}

func editPage(writer http.ResponseWriter, request *http.Request) {
	title, _ := utils.GetPageTitle(writer, request)
	page, err := data.LoadPage(title)

	if err != nil {
		data.CreateTempPageAndSave(title)
		editPage(writer, request)
	} else {
		err := templates.ExecuteTemplate(writer, "edit.html", page)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func EditPageHandler() {
	http.HandleFunc("/edit/", editPage)
	fmt.Println("Starting editing handler")
}

func saveUpdatedPage(writer http.ResponseWriter, request *http.Request) {
	title, _ := utils.GetPageTitle(writer, request)
	updatedBody := request.FormValue("body")

	updatePage := data.Page{Title: title, Body: []byte(updatedBody)}
	updatePage.Save()
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}

func SavePageHandler() {
	http.HandleFunc("/save/", saveUpdatedPage)
	fmt.Println("Starting save page handler")
}

func ListenAndServe() {
	log.Fatalln(http.ListenAndServe(port, nil))
}
