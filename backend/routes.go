package main

import (
	"log"
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

	log.Println("Listening...")

	fs := http.FileServer(http.Dir("../frontend"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.Handle("/", fs)

	return router
}

var routes = Routes{
	Route{
		"FilesIndex",
		"GET",
		"/files",
		FilesIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/files/{fileID}",
		FileShow,
	},
	{
		"FileUpload",
		"POST",
		"/upload",
		FileUpload,
	},
	{
		"FileUpload",
		"GET",
		"/upload",
		FileUpload,
	},
}
