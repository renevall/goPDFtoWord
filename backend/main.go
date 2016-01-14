package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	//router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
