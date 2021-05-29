/**
 *
 **/
package main

import (
	"net/http"

	"github.com/ipenguin/continuous-bible-plans/internal/app"
)

func main() {

	http.HandleFunc("/", app.Routes)
	http.ListenAndServe(":8080", nil)
}
