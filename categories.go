package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
 * Category... - a post struct for all my posts
 */
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AllCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All Categories!")
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]
	fmt.Fprintln(w, "Get Category:", categoryID)
}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert Category!")
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Category!")
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Category!")
}
