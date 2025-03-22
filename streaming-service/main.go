package main

import (
	"fmt"
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Streaming content..."))
}

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("Auth Service is running on port 8083/stream...")
	http.ListenAndServe(":8083", nil)
}

