package handlers

import (
	"encoding/json"
	"net/http"

	"supply-chain-insight/backend/internal/services"

	"gorm.io/gorm"
)

func RegisterUserRoutes(mux *http.ServeMux, db *gorm.DB) {
	userService := services.NewUserService(db)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		users, err := userService.GetUsers()
		if err != nil {
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})
}
