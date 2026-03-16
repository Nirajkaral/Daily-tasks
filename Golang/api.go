package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	msg := Message{Text: "Hello from API"}

	json.NewEncoder(w).Encode(msg)
}

func api() {

	http.HandleFunc("/api", apiHandler)

	http.ListenAndServe(":8080", nil)
}
