package main

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TutorialEdge REST API: v0.0.1")
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	// All Tags Routes
	Route{"Alltags", "GET", "/tags", AllTags},
	Route{"Gettag", "GET", "/tag/{id}", GetTag},
	Route{"NewTag", "POST", "/tag", InsertTag},
	Route{"EditTag", "POST", "/tag/{id}", EditTag},
	Route{"DeleteTag", "DELETE", "/tag/{id}", DeleteTag},
	// All Blog Posts Routes
	// ...
	Route{"AllPosts", "GET", "/posts", AllPosts},
	Route{"GetPost", "GET", "/post/{id}", GetPost},
	Route{"NewPost", "POST", "/post", InsertPost},
	Route{"EditPost", "POST", "/post/{id}", EditPost},
	Route{"DeletePost", "DELETE", "/post/{id}", AllPosts},
	// All Tutorials Routes
	Route{"AllTutorials", "GET", "/tutorials", AllTutorials},
	Route{"GetTutorial", "GET", "/tutorial/{id}", GetTutorial},
	Route{"NewTutorial", "POST", "/tutorial", InsertTutorial},
	Route{"EditTutorial", "POST", "/tutorial/{id}", EditTutorial},
	Route{"DeleteTutorial", "DELETE", "/tutorial/{id}", AllTutorials},
	// All Categories
	Route{"AllCategories", "GET", "/categories", AllCategories},
	Route{"GetCategory", "GET", "/category/{id}", GetCategory},
	Route{"NewCategory", "POST", "/category", InsertCategory},
	Route{"EditCategory", "POST", "/category/{id}", EditCategory},
	Route{"DeleteCategory", "DELETE", "/category/{id}", AllCategories},
	// All Categories
	Route{"AllCourses", "GET", "/courses", AllCourses},
	// Route{"GetCategory", "GET", "/category/{id}", GetCategory},
	// Route{"NewCategory", "POST", "/category", InsertCategory},
	// Route{"EditCategory", "POST", "/category/{id}", EditCategory},
	// Route{"DeleteCategory", "DELETE", "/category/{id}", AllCategories},
}
