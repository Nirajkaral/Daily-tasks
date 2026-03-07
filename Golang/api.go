package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

// Renamed to apiHandler - handles the /api route
func apiHandler(w http.ResponseWriter, r *http.Request) {

	msg := Message{Text: "Hello from API"}

	json.NewEncoder(w).Encode(msg)
}

// Renamed to main - starts the server
func structs() {

	http.HandleFunc("/api", apiHandler)

	http.ListenAndServe(":8080", nil)
}