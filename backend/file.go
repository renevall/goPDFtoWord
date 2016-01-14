package main

import "time"

type File struct {
	Name     string    `json:"name"`
	Size     string    `json:"size"`
	Uploaded time.Time `json:"uploaded"`
}

type Files []File
