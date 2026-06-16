package handlers

import (
	"encoding/json"
	"net/http"
	"supply-chain-insight/backend/internal/services"

	"gorm.io/gorm"
)

func RegisterProductRoutes(mux *http.ServeMux, db *gorm.DB) {
	productService := services.NewProductService(db)
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		products, err := productService.GetProducts()
		if err != nil {
			http.Error(w, "failed to get products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})
}
