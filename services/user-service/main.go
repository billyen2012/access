package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Service   string `json:"service"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Timestamp string `json:"timestamp"`
	Data      any    `json:"data"`
}

var users = []map[string]any{
	{"id": 1, "name": "Alice", "email": "alice@example.com", "role": "admin"},
	{"id": 2, "name": "Bob", "email": "bob@example.com", "role": "user"},
	{"id": 3, "name": "Charlie", "email": "charlie@example.com", "role": "user"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Service:   "user-service",
		Path:      r.URL.Path,
		Method:    r.Method,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      users,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy", "service": "user-service"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handler)
	log.Println("user-service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
