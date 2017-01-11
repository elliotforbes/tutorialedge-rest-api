package main

import (
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
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"description"`
	Body      string `json:"body"`
	IsLive    int    `json:"isLive"`
	Author    string `json:"author"`
	Slug      string `json:"slug"`
	ImagePath string `json:"image_path"`
}

func AllTutorials(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var tutorials []Tutorial
	results, err := db.Query("SELECT id, title, description, body, slug, isLive, author, image_path FROM lessons")

	for results.Next() {
		var tutorial Tutorial
		err = results.Scan(&tutorial.ID, &tutorial.Title, &tutorial.Desc, &tutorial.Body, &tutorial.Slug, &tutorial.IsLive, &tutorial.Author, &tutorial.ImagePath)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select all from tutorials"})
		}
		tutorials = append(tutorials, tutorial)
	}

	json.NewEncoder(w).Encode(tutorials)
}

func GetTutorial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tutorialID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}
	// Open up our database connection.
	db := connect()
	defer db.Close()

	var tutorial Tutorial
	// Execute the query
	err = db.QueryRow("SELECT id, title, description, body, slug, isLive, author, image_path FROM lessons where id = ?", tutorialID).Scan(&tutorial.ID, &tutorial.Title, &tutorial.Desc, &tutorial.Body, &tutorial.Slug, &tutorial.IsLive, &tutorial.Author, &tutorial.ImagePath)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		fmt.Fprintln(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(tutorial)
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
