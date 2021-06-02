package indexpage

import (
	"log"
	"net/http"

	"github.com/ipenguin/continuous-bible-plans/internal/page"
)

const (
	// Define the URL for this page
	_pageName = "/"
)

type TestData struct {
	Key string
}

// Return the URL for this page for use in routing and other external functions
func PageNames() []string {
	return []string{_pageName}
}

// Deliver the data for this page to the client
func GeneratePage(response http.ResponseWriter, request *http.Request) {

	log.Println(request.URL.Path)
	page.GeneratePage(response, "assets/templates/index.html")
}
