package page

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	name           string
	page_generator func(response http.ResponseWriter, request *http.Request)
}

var all_pages []Page

func GeneratePage(response http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles(templateName)
	err := parsedTemplate.Execute(response, "")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func AddPages(p Page) {
	all_pages = append(all_pages, p)
}

func AllPages() []Page {
	return all_pages
}
