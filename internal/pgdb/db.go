package pgdb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Pgdb struct {
	Address  string
	Port     int
	Username string
	Password string
	DB       string
}

func Init(dbinfo Pgdb) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbinfo.Address, dbinfo.Port, dbinfo.Username, dbinfo.Password, dbinfo.DB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// APPLY SCHEMA
	sqlFile, err := os.ReadFile("internal/pgdb/query/schema.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return nil, err
	}
	fmt.Println("Schema applied")
	return db, nil
}
