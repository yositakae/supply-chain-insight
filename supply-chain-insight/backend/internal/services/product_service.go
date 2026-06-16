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
func (s *ProductService) GetProductByID(id string) (*models.ProductResponse, error) {
	var product models.ProductResponse

	// Product IDs are varchar in the database, so keep the route id as a string.
	result := s.db.
		Table("products").
		Select("products.id, products.product_name, category.category AS category, products.price, products.stock").
		Joins("LEFT JOIN category ON category.id = products.category_id").
		Where("products.id = ?", id).
		Scan(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &product, nil
}
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.db.Create(product).Error
}
func (s *ProductService) EditProduct(product *models.Product) error {
	return s.db.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(product).Error
}
