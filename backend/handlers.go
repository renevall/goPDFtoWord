package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func FilesIndex(w http.ResponseWriter, r *http.Request) {

	files := Files{
		File{Name: "Algo"},
		File{Name: "Fuiste tu"},
	}

	if err := json.NewEncoder(w).Encode(files); err != nil {
		panic(err)
	}
}

func FileShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileID"]
	fmt.Fprintln(w, "Todo show:", fileID)
}
