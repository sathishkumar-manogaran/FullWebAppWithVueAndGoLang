package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Host name for this local app
	Host = "localhost"
	// Port number for this service
	Port = "8080"
)

// change this func name before running
func main1() {
	fmt.Printf("test")
	http.HandleFunc("/", home)
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Occoured : ", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample Http server, up and running")
}

// Simple default net/http program, this is the default package to run simple server in go
// Only one problem is, doing dynamic routing is difficult in default net/http
// So we will use gorilla/mux to handle the problem
