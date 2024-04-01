package main

import (
	"fmt"
	server "goplayground/api"
	"goplayground/internal/pgdb"
)

func main() {
	dbinfo := pgdb.Pgdb{
		Address:  "localhost",
		Port:     5432,
		Username: "user",
		Password: "pass",
		DB:       "db",
	}
	db, err := pgdb.Init(dbinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to database")

	defer db.Close()

	//	user := pgdb.User{
	//Username: "testuser",
	//Password: "testpassword",
	//Email:    "testuser@example.com",
	//}

	//err = pgdb.CreateUser(db, user)
	//if err != nil {
	//log.Fatal(err)
	//}

	API := server.App{
		DB: db,
	}

	err = server.StartHTTPServer(API)
	if err != nil {
		fmt.Println(err)
	}
}
