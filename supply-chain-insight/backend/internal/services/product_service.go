package services

import (
	"supply-chain-insight/backend/internal/models"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}
func (s *ProductService) GetProducts() ([]models.ProductResponse, error) {
	var products []models.ProductResponse
	if err := s.db.
		Table("products").
		Select("products.id, products.product_name, category.category AS category, products.price, products.stock").
		Joins("LEFT JOIN category ON category.id = products.category_id").
		Order("products.id").
		Scan(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
