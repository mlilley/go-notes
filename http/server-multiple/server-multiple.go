package main

import (
	"fmt"
	"net/http"
)

// 1. create a ServeMux per server for separate routing
// 2. pass mux to http.ListenAndServe(port, mux) for each server instance
//    - one server should block the main go routine to keep the app alive
//    - the other started in a separate go routine

func getRootHandler1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Srv1 root!")
}

func getRootHandler2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Srv2 root!")
}

func main() {
	srv1 := http.NewServeMux()
	srv2 := http.NewServeMux()

	srv1.HandleFunc("/", getRootHandler1)
	srv2.HandleFunc("/", getRootHandler2)

	go func() {
		http.ListenAndServe(":8002", srv2)
	}()

	http.ListenAndServe(":8001", srv1)
}
