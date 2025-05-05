package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/user/")
		if id == "" {
			http.Error(w, "Missing user ID", http.StatusBadRequest)
			return
		}

		// Mock user data
		user := User{
			ID:   id,
			Name: fmt.Sprintf("User %s", id),
			Age:  25,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	port := "8081"
	log.Printf("Starting user-service on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
