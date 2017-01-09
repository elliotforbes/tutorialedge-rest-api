package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
 * Post... - a post struct for all my posts
 */
type Post struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All Posts!")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]
	fmt.Fprintln(w, "Get Post:", postID)
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert Post!")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Post!")
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Post!")
}
