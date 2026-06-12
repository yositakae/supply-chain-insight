package handlers

import (
	"encoding/json"
	"net/http"

	"supply-chain-insight/backend/internal/services"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	userService := services.NewUserService()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userService.GetUsers())
	})
}
