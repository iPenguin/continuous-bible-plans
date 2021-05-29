package app

import (
	"fmt"
	"net/http"
)

func Routes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Something to read")
}
