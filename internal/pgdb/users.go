package pgdb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt sql.NullTime
}

// List users
func ListUsers(db *sql.DB) ([]User, error) {
	query := `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create a new user
func CreateUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, user.Username, user.Password, user.Email)
	return err
}

// Get a user by ID
func GetUser(db *sql.DB, id int) (User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	return user, err
}

// Update a user
func UpdateUser(db *sql.DB, user User) error {
	query := `UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4`
	_, err := db.Exec(query, user.Username, user.Password, user.Email, user.ID)
	return err
}

// Delete a user
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
