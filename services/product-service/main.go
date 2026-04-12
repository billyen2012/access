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

var products = []map[string]any{
	{"id": 1, "name": "Laptop", "price": 999.99, "stock": 50, "category": "electronics"},
	{"id": 2, "name": "Phone", "price": 699.99, "stock": 120, "category": "electronics"},
	{"id": 3, "name": "Headphones", "price": 149.99, "stock": 200, "category": "accessories"},
	{"id": 4, "name": "Keyboard", "price": 79.99, "stock": 300, "category": "accessories"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Service:   "product-service",
		Path:      r.URL.Path,
		Method:    r.Method,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      products,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy", "service": "product-service"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handler)
	log.Println("product-service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
