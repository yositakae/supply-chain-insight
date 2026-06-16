package models

type Product struct {
	ID          string  `gorm:"primaryKey" json:"id"`
	ProductName string  `gorm:"column:product_name" json:"product_name"`
	CategoryID  int     `gorm:"column:category_id" json:"category_id"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
type ProductResponse struct {
	ID          string  `gorm:"primaryKey" json:"id"`
	ProductName string  `gorm:"column:product_name" json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
