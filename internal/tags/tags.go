package tags

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipenguin/continuous-bible-plans/internal/page"
)

const (
	// Define the URL for this page
	_tags       = "/tags"
	_single_tag = "/tags/{tag_name}"
)

type TestData struct {
	Key string
}

// Return the URL for this page for use in routing and other external functions
func PageNames() []string {
	return []string{_tags, _single_tag}
}

// Deliver the data for this page to the client
func GeneratePage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tag_name := vars["tag_name"]

	switch tag_name {
	case "":
		page.GeneratePage(response, "assets/templates/tags.html")
	default:
		log.Println("Tag:", tag_name)
	}
}
