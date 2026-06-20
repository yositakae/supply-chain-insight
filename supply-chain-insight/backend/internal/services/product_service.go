package services

import (
	"strings"

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
	product.ProductName = strings.TrimSpace(product.ProductName)

	// แก้ให้เพิ่มข้อมูลได้: ถ้าไม่ได้ส่ง id มา ให้สร้าง id ถัดไปจากข้อมูลเดิมก่อน insert
	if product.ID == "" {
		nextID, err := s.nextProductID()
		if err != nil {
			return err
		}
		product.ID = nextID
	}

	return s.db.Table("products").Create(product).Error
}

func (s *ProductService) nextProductID() (string, error) {
	var nextID string

	// แก้ให้เพิ่มข้อมูลได้: id ในตารางเป็น varchar จึงแปลงเฉพาะ id ที่เป็นตัวเลขเพื่อหาเลขถัดไป
	err := s.db.
		Raw("SELECT (COALESCE(MAX(id::integer), 0) + 1)::text FROM products WHERE id ~ ?", "^[0-9]+$").
		Scan(&nextID).Error

	return nextID, err
}
func (s *ProductService) EditProduct(product *models.Product) error {
	result := s.db.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(product)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (s *ProductService) DeleteProduct(id string) error {
	result := s.db.Delete(&models.Product{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
