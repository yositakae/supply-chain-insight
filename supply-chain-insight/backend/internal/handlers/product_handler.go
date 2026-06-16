package handlers

import (
	"encoding/json"
	"net/http"
	"supply-chain-insight/backend/internal/models"
	"supply-chain-insight/backend/internal/services"

	"gorm.io/gorm"
)

func RegisterProductRoutes(mux *http.ServeMux, db *gorm.DB) {
	productService := services.NewProductService(db)
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			products, err := productService.GetProducts()
			if err != nil {
				http.Error(w, "failed to get products", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)

		case http.MethodPost:
			var product models.Product
			if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := productService.CreateProduct(&product); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
