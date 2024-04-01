package server

import (
	"database/sql"
	v1 "goplayground/api/v1"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func StartHTTPServer(app App) error {
	mux := http.NewServeMux()

	// Swagger Components
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./api/swagger"))))
	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/swagger/oapi.json")
	})

	apiPrefix := "/api/v1"

	// Applications
	applicationPrefix := apiPrefix + "/applications/"
	mux.HandleFunc(applicationPrefix, v1.ApplicationHandler)

	return http.ListenAndServe(":8080", mux)
}
