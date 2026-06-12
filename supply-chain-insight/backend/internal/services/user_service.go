package services

import "supply-chain-insight/backend/internal/models"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers() []models.User {
	return []models.User{
		{
			ID:    1,
			Name:  "Demo User",
			Email: "demo@example.com",
		},
	}
}
