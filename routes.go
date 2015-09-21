package main

import (
	"net/http"
)

//Route strucuture that define our routing
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice of route
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"OPTIONS",
		"/todos",
		AllowAcces,
	},
	Route{
		"Index",
		"OPTIONS",
		"/todos/{todoId}",
		AllowAcces,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Index",
		"GET",
		"/todos",
		TodoShow,
	},
	Route{
		"TodoIndex",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"PUT",
		"/todos/{todoId}",
		TodoUpdate,
	},
	Route{
		"TodoShow",
		"DELETE",
		"/todos/{todoId}",
		TodoDelete,
	},
}
