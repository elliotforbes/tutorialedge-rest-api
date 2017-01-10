package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
 * Tag... - standard tag used in frontend
 */
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AllTags(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var tags []Tag
	results, err := db.Query("SELECT id, name FROM tags")

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			log.Print(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select tag from database"})
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}
	db := connect()
	defer db.Close()

	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM tags where id = ?", tagID).Scan(&tag.ID, &tag.Name)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select tag from database"})
	}

	json.NewEncoder(w).Encode(tag)
}

func InsertTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tag Tag
	err := decoder.Decode(&tag)

	if err != nil {
		log.Print(err.Error())
	}
	db := connect()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO tags (name) values (?)")
	res, err := stmt.Exec(tag.Name)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select tag from database"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted Tag Into the Database", Body: strconv.Itoa(int(id))})
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	tagID := vars["id"]
	ID, _ := strconv.Atoi(tagID)

	stmt, err := db.Prepare("DELETE FROM tags where id = ?")
	if err != nil {
		log.Print(err.Error())
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to delete tag from database"})
	}
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Deleted Tag from the Database"})

}

func EditTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Tag!")
}
