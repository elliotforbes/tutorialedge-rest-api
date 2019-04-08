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
 * Post... - a post struct for all my posts
 */
type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"description"`
	Body   string `json:"body"`
	IsLive int    `json:"isLive"`
	Author string `json:"author"`
	Slug   string `json:"slug"`
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var posts []Post
	results, err := db.Query("SELECT id, title, description, body, slug, isLive, author FROM posts")

	for results.Next() {
		var post Post
		err = results.Scan(&post.ID, &post.Title, &post.Desc, &post.Body, &post.Slug, &post.IsLive, &post.Author)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select all from posts"})
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintln(w, "Not a Valid id")
	}

	var post Post
	// Execute the query
	err = db.QueryRow("SELECT id, title, description, body, slug, isLive, author FROM posts where id = ?", postID).Scan(&post.ID, &post.Title, &post.Desc, &post.Body, &post.Slug, &post.IsLive, &post.Author)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to select tag from database"})
	}

	json.NewEncoder(w).Encode(post)
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var post Post
	err := decoder.Decode(&post)

	if err != nil {
		log.Print(err.Error())
	}

	stmt, _ := db.Prepare("INSERT INTO posts (title, description, body, isLive, author, slug, created_at, updated_at, published_at) values (?,?,?,?,?,?, NOW(), NOW(), NOW()) ")
	res, err := stmt.Exec(post.Title, post.Desc, post.Body, post.IsLive, post.Author, post.Slug)
	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert post into database"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id"})
	}

	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted Post Into the Database", Body: strconv.Itoa(int(id))})

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	vars := mux.Vars(r)
	postID := vars["id"]
	ID, _ := strconv.Atoi(postID)

	stmt, err := db.Prepare("DELETE FROM posts where id = ?")
	if err != nil {
		log.Print(err.Error())
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to delete post from database"})
	}
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Deleted Post from the Database"})

}

func EditPost(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var post Post
	err := decoder.Decode(&post)

	vars := mux.Vars(r)
	postID := vars["id"]
	ID, _ := strconv.Atoi(postID)

	stmt, _ := db.Prepare("UPDATE posts SET title = ?, description = ?, body = ?, author = ?, updated_at = NOW() WHERE id = ?")

	_, err = stmt.Exec(post.Title, post.Desc, post.Body, post.Author, ID)

	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to Update Post in the Database"})
	}
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Update Post in the Database"})
}
