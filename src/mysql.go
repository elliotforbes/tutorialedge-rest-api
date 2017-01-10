package main

import (
	"database/sql"
	"log"
)

type DB struct {
	connection *DB
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:charlie1@tcp(127.0.0.1:3306)/tuts")

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	return db
}
