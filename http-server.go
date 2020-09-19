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

func main() {
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
