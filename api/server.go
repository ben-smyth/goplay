package server

import (
	"database/sql"
	"fmt"
	v1 "goplayground/api/v1"
	"goplayground/web"
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

	// Auth
	// mux.HandleFunc("/login", auth.LoginHandler)
	// mux.HandleFunc("/oidc/callback", auth.OidcCallback)

	apiPrefix := "/api/v1"

	// Applications
	applicationPrefix := apiPrefix + "/applications/"
	mux.HandleFunc(applicationPrefix, v1.ApplicationHandler)
	mux.HandleFunc(applicationPrefix+"{name}", v1.ApplicationByName)

	// UI
	templates := web.PopulateTemplates()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := templates.Lookup("index")
		if tmpl == nil {
			fmt.Println("template not found")
			return
		}
		err := tmpl.Execute(w, "")
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	return http.ListenAndServe(":8080", mux)
}
