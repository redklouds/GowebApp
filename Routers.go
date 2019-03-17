package main

/*
Routers.go
This file is to keep the structure clean
This file houses all the routers of the application,
meaning that each time a request comes in we rout to the correct location

use this class by calling the NewRouter Function which returns the
initalized router
*/

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	//create the router
	router := mux.NewRouter().StrictSlash(true)

	//router definitions

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/todos", TodoIndex).Methods("GET")
	router.HandleFunc("/todos/{todoId}", TodoShow).Methods("GET")

	//post router
	router.HandleFunc("/todos", TodoCreate).Methods("POST")

	return router
}

//The officeal way of doing things is below
//below we define a array that will instantiate the handler and their endpoints

/*
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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
*/
