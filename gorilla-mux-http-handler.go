package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Host1 name for this local app
	Host1 = "localhost"
	// Port1 number for this service
	Port1 = "8080"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", GetRequest).Methods("GET")
	http.ListenAndServe(Host1+":"+Port1, router)
}

// GetRequest Basic
var GetRequest = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get request received success")
	},
)
