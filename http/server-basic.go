package main

// Note:
//  - cannot have multiple servers per process with this method
//  - url matching behavior may be surprising:
//    - "/" is a catch all for everything not matched by other handlers
//    - "/x" matches "/x" and "/xyz"

import (
	"fmt"
	"net/http"
)

func getRootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "I am root!")
}

func getNotesHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Notes!")
}

func main() {
	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/notes", getNotesHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Unalbe to start server")
	}
}
