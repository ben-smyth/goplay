package server

import (
	"encoding/json"
	"goplayground/internal/pgdb"
	"net/http"
	"strconv"
)

// swagger:response errorResponse
type errorResponse struct {
	Error string `json:"error"`
}

func jsonError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse{Error: error})
}

// listUsers lists all users.
// swagger:route GET /users users listUsers
//
// Lists all users.
//
// Produces:
// - application/json
//
// Responses:
//
//	200:
//	500: errorResponse
func (app *App) listUsers(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := pgdb.ListUsers(app.DB)
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dbUsers)
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user pgdb.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := pgdb.CreateUser(app.DB, user)
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := pgdb.GetUser(app.DB, id)
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *App) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var user pgdb.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = id

	err := pgdb.UpdateUser(app.DB, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *App) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	err := pgdb.DeleteUser(app.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
