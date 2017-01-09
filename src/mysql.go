package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getTag() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:charlie1@tcp(127.0.0.1:3306)/tuts")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(tag.ID)
	log.Println(tag.Name)

}
