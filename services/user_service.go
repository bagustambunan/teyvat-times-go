package services

import (
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type UserService interface {
	UpdateUserCoins(user *models.User, coins int) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

type USConfig struct {
	UserRepository repositories.UserRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository: c.UserRepository,
	}
}

func (serv *userService) UpdateUserCoins(user *models.User, coins int) (*models.User, error) {
	return serv.userRepository.UpdateCoins(user, coins)
}
