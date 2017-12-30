package main

import (
	"net/http"
)

// Serves a file list view and file contents

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
