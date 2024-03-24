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
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/swagger/swagger.json")
	})

	// API ENDPOINTS
	// Users
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.listUsers(w, r)
		case http.MethodPost:
			app.createUser(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.getUser(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	return http.ListenAndServe(":8080", mux)
}
