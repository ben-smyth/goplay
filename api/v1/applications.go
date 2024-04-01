package v1

import (
	"net/http"
)

func ApplicationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	default:

	}
}

func ApplicationByName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodDelete:

	case http.MethodPatch:


	default:
}
