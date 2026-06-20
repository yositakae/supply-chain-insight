package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"supply-chain-insight/backend/internal/models"
	"supply-chain-insight/backend/internal/services"

	"gorm.io/gorm"
)

func RegisterProductRoutes(mux *http.ServeMux, db *gorm.DB) {
	productService := services.NewProductService(db)

	// ส่วนแก้ไขสินค้า: รับ id จาก URL เช่น /products/222 แล้วอัปเดตข้อมูลจาก JSON body.
	updateProduct := func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// แก้ตรงนี้: ตรวจว่า id ใน body ต้องตรงกับ id ใน URL
		// เช่น PATCH /products/222 แต่ body เป็น "id": "223" จะคืน 400 Bad Request.
		// ถ้าไม่ส่ง id ใน body จะใช้ id จาก URL แทน.
		urlID := r.PathValue("id")
		if product.ID != "" && product.ID != urlID {
			http.Error(w, "product id in body does not match URL", http.StatusBadRequest)
			return
		}

		product.ID = urlID
		if err := productService.EditProduct(&product); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "product not found", http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}

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

			// แก้ให้เพิ่มข้อมูลได้: ตรวจ field ที่จำเป็นก่อนส่งเข้า service เพื่อให้รู้สาเหตุชัดเจน
			if strings.TrimSpace(product.ProductName) == "" {
				http.Error(w, "product_name is required", http.StatusBadRequest)
				return
			}
			if product.CategoryID <= 0 {
				http.Error(w, "category_id is required", http.StatusBadRequest)
				return
			}

			if err := productService.CreateProduct(&product); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// แก้ให้เพิ่มข้อมูลได้: คืนข้อมูลที่สร้างกลับไป รวมถึง id ที่ service สร้างให้อัตโนมัติ
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// ส่วนดูสินค้าตาม id: ใช้ GET /products/{id}
	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		product, err := productService.GetProductByID(r.PathValue("id"))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "product not found", http.StatusNotFound)
				return
			}

			http.Error(w, "failed to get product", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	})

	// ส่วน route สำหรับแก้ไขสินค้า: Postman ยิง POST หรือ PATCH มาที่ /products/{id} ได้.
	mux.HandleFunc("POST /products/{id}", updateProduct)
	mux.HandleFunc("PATCH /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		if err := productService.DeleteProduct(id); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "product not found", http.StatusNotFound)
				return
			}

			http.Error(w, "failed to delete product", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
