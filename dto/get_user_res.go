package dto

import "final-project-backend/models"

type GetUserRes struct {
	ID         int             `json:"userID"`
	Role       *models.Role    `json:"role"`
	Username   string          `json:"username"`
	Email      string          `json:"email"`
	Name       string          `json:"name"`
	Phone      string          `json:"phone"`
	Address    *models.Address `json:"address"`
	ProfilePic *models.Image   `json:"profilePic"`
}

func (_ *GetUserRes) FromUser(u *models.User) *GetUserRes {
	return &GetUserRes{
		ID:         u.ID,
		Role:       u.Role,
		Username:   u.Username,
		Email:      u.Email,
		Name:       u.Name,
		Phone:      u.Phone,
		Address:    u.Address,
		ProfilePic: u.ProfilePic,
	}
}
