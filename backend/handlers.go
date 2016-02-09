package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	err  string
	code bool
}

type Response2 struct {
	status        bool
	originalName  string
	generatedName string
}

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

func FileUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("METHOD IS " + r.Method + " AND CONTENT-TYPE IS " + r.Header.Get("Content-Type"))

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		json.NewEncoder(w).Encode(Response{err.Error(), true})
		fmt.Println(Response{err.Error(), true})
		return
	}
	defer file.Close()
	// fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		json.NewEncoder(w).Encode(Response{err.Error(), true})
		fmt.Println(Response{err.Error(), true})

		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		json.NewEncoder(w).Encode(Response{err.Error(), true})
		return
	}
	fmt.Println(Response{"File '" + handler.Filename + "' submited successfully", false})
	json.NewEncoder(w).Encode(Response2{true, handler.Filename + "Server", handler.Filename})
}
