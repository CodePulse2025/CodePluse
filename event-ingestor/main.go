package main

import (
	"fmt"
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// For now, just read and print the body
	defer r.Body.Close()
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	log.Printf("Received webhook: %s", string(body))

	// (Later: Parse the body, send to gRPC client)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := 9000
	log.Printf("Event Ingestor listening on port %d...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
