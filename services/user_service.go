package services

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type UserService interface {
	UpdateUserMora(user *models.User, mora int) (*models.User, error)
	GetUser(user *models.User) (*dto.GetUserRes, error)
	GetUserDownLines(user *models.User) ([]*models.User, error)
	GetUserReferral(user *models.User) (*models.UserReferral, error)
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

func (serv *userService) UpdateUserMora(user *models.User, mora int) (*models.User, error) {
	return serv.userRepository.UpdateMora(user, mora)
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

func (serv *userService) GetUserReferral(user *models.User) (*models.UserReferral, error) {
	return serv.userRepository.FindUserReferral(user)
}
