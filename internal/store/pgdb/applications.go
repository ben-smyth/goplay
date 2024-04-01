package pgdb

import (
	"database/sql"
	"errors"
)

type Application struct {
	ID          int
	Name        string
	Description string
	Public      bool
	CreatedAt   sql.NullTime
}

func GetApplicationByName(name string, db *sql.DB) (Application, error) {
	var app Application

	query := `SELECT id,name,description,public,created_at FROM applications WHERE name = $1`
	row := db.QueryRow(query, name)

	err := row.Scan(&app.ID, &app.Name, &app.Description, &app.Public, &app.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return app, NewSQLNotFound("application", name)
		}
		return app, err
	}
	return app, nil
}

func DeleteApplicationByName(name string, db *sql.DB) error {
	query := `DELETE FROM applications WHERE name = $1`
	_, err := db.Exec(query, name)
	if err != nil {
		return NewSQLDeleteFailed("application", name, err)
	}
	return nil
}

func UpdateApplicationByName(name string, app Application, db *sql.DB) error {
	query := `UPDATE applications SET name = $1, description = $2, public = $3 WHERE name = $1`
	_, err := db.Exec(query, app.Name, app.Description, app.Public)
	if err != nil {
		return NewSQLUpdateFailed("application", name, err)
	}
	return nil
}

func CreateApplication(app Application, db *sql.DB) error {
	query := `INSERT INTO applications (name,description,public) VALUES ($1,$2,$3)`
	_, err := db.Exec(query, app.Name, app.Description, app.Public)
	if err != nil {
		return NewSQLCreateFailed("application", app.Name, err)
	}
	return nil
}

func ListApplications(db *sql.DB) ([]Application, error) {
	var apps []Application

	query := `SELECT id, name, description, public, created_at FROM applications`
	rows, err := db.Query(query)
	if err != nil {
		return nil, NewSQLListFailed("application", err)
	}
	defer rows.Close()

	for rows.Next() {
		var app Application
		if err := rows.Scan(&app.ID, &app.Name, &app.Description, &app.Public, &app.CreatedAt); err != nil {
			return nil, NewSQLListFailed("application", err)
		}
		apps = append(apps, app)
	}
	if err := rows.Err(); err != nil {
		return nil, NewSQLListFailed("application", err)
	}

	return apps, nil
}
