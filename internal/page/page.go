package page

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Page interface {
	GeneratePage(router mux.Router)
	PageNames() string
}

func GeneratePage(response http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles(templateName)
	err := parsedTemplate.Execute(response, "")
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
