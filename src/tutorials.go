package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
 * Tutorial... - standard tag used in frontend
 */
type Tutorial struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AllTutorials(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All Tags!")
}

func GetTutorial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:charlie1@tcp(127.0.0.1:3306)/tuts")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM tags where id = ?", tagID).Scan(&tag.ID, &tag.Name)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		fmt.Fprintln(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(tag)
}

func InsertTutorial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert Tutorial!")
}

func DeleteTutorial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Tutorial!")
}

func EditTutorial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Tutorial!")
}
