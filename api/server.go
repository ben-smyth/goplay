package server

import (
	"database/sql"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func StartHTTPServer(app App) error {
	mux := http.NewServeMux()

	// Swagger Components
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./api/swagger"))))
	mux.HandleFunc("/oapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/swagger/oapi.json")
	})

	return http.ListenAndServe(":8080", mux)
}
