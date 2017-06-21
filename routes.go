package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Resize",
		"GET",
		"/s/{params}/{url}.{type}",
		Resize,
	},
	Route{
		"Crop",
		"GET",
		"/c/{params}/{url}.{type}",
		Crop,
	},
}
