package models

type Product struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"column:product_name" json:"product_name"`
	CategoryID  int    `gorm:"column:category_id" json:"category_id"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
}
