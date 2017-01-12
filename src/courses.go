package main

import (
	"encoding/json"
	"net/http"
)

/*
 * Course... a struct to represent a course in the database
 */
type Course struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func AllCourses(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var courses []Course
	results, err := db.Query("SELECT id, title FROM courses")

	for results.Next() {
		var course Course
		err = results.Scan(&course.ID, &course.Title)
		if err != nil {
			json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to retrieve all courses"})
		}
		courses = append(courses, course)
	}

	json.NewEncoder(w).Encode(courses)

}
