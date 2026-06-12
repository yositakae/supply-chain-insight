package main

import (
	"log"
	"net/http"

	"supply-chain-insight/backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.RegisterUserRoutes(mux)

	port := "8080"

	addr := ":" + port
	log.Printf("backend listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
