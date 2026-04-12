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

var orders = []map[string]any{
	{"id": 1001, "userId": 1, "product": "Laptop", "amount": 999.99, "status": "shipped"},
	{"id": 1002, "userId": 2, "product": "Phone", "amount": 699.99, "status": "pending"},
	{"id": 1003, "userId": 1, "product": "Headphones", "amount": 149.99, "status": "delivered"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Service:   "order-service",
		Path:      r.URL.Path,
		Method:    r.Method,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      orders,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy", "service": "order-service"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handler)
	log.Println("order-service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
