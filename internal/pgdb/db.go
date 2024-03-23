package pgdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Pgdb struct {
	Address  string
	Port     int
	Username string
	Password string
	DB       string
}

func Init(dbinfo Pgdb) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbinfo.Address, dbinfo.Port, dbinfo.Username, dbinfo.Password, dbinfo.DB)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
