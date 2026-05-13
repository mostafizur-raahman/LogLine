package main

import (
	"log"
	"logline/internal/server"
	"net/http"
)

func main() {
	srv := server.New()

	addr := ":4000"

	log.Printf("loging starting on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, srv))
}
