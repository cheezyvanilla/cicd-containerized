package main

// create ping rest api using http
// create ping service

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// activity log when api is hit
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Ping request received")
		fmt.Fprintf(w, "pong")
	})

	logger.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
