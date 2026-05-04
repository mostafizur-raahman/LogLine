package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
func handleIngest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
	w.WriteHeader(http.StatusCreated)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("POST /ingest", handleIngest)

	addr := ":4000"

	log.Printf("logline starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
