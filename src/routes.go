package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TutorialEdge REST API: v0.0.1!")
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},

	// All Tags Routes
	Route{"Alltags", "GET", "/tags", AllTags},
	Route{"Gettag", "GET", "/tag/{id}", GetTag},
	Route{"NewTag", "POST", "/tag", InsertTag},
	Route{"EditTag", "POST", "/tag/{id}", EditTag},
	Route{"DeleteTag", "DELETE", "/tag/{id}", DeleteTag},
	// All Posts Routes
	// ...
	Route{"AllPosts", "GET", "/posts", AllPosts},
}
