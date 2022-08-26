package services

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type UserService interface {
	UpdateUserCoins(user *models.User, coins int) (*models.User, error)
	GetUser(user *models.User) (*dto.GetUserRes, error)
	GetUserDownLines(user *models.User) ([]*models.User, error)
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

func (serv *userService) GetUser(user *models.User) (*dto.GetUserRes, error) {
	fetchedUser, fetchErr := serv.userRepository.FindUser(user)
	if fetchErr != nil {
		return nil, fetchErr
	}
	return new(dto.GetUserRes).FromUser(fetchedUser), nil
}

func (serv *userService) GetUserDownLines(user *models.User) ([]*models.User, error) {
	return serv.userRepository.GetUserDownLines(user)
}
