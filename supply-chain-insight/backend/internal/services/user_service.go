package services

import (
	"supply-chain-insight/backend/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Order("id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
