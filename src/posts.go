package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
 * Post... - a post struct for all my posts
 */
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var posts []Post
	results, err := db.Query("SELECT id, title FROM posts")

	for results.Next() {
		var post Post
		err = results.Scan(&post.ID, &post.Title)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Failed to select all from posts"})
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	postID := vars["id"]
	fmt.Fprintln(w, "Get Post:", postID)
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	fmt.Fprintln(w, "Insert Post!")
	decoder := json.NewDecoder(r.Body)
	var post Post
	err := decoder.Decode(&post)

	if err != nil {
		log.Print(err.Error())
	}

	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Post!")
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Edit Post!")
}
