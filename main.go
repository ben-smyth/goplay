package main

import "goplayground/internal/pgdb"

func main() {
	db := pgdb.Pgdb{
		Address:  "localhost",
		Port:     5432,
		Username: "citizix_user",
		Password: "S3cret",
		DB:       "citizix_db",
	}
	err := pgdb.Init(db)
	if err != nil {
		return
	}
}
