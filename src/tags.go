package main

import (
	"fmt"
	"net/http"

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
	fmt.Fprintln(w, "All Tags!")
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagID := vars["id"]
	fmt.Fprintln(w, "Get Tag:", tagID)
}

func InsertTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert Tag!")
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Tag!")
}

func EditTag(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Tag!")
}
