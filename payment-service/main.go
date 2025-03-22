package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type PaymentRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)

	// Simulate payment processing
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful"))
}

func main() {
	http.HandleFunc("/pay", paymentHandler)
	fmt.Println("Auth Service is running on port 8082/pay...")
	http.ListenAndServe(":8082", nil)
}

