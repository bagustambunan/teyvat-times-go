package dto

import "final-project-backend/models"

type SignUpRes struct {
	ID      int             `json:"userID"`
	Email   string          `json:"email"`
	Name    string          `json:"name"`
	Phone   string          `json:"phone"`
	Address *models.Address `json:"address"`
}

func (_ *SignUpRes) FromUser(u *models.User) *SignUpRes {
	return &SignUpRes{
		ID:      u.ID,
		Email:   u.Email,
		Name:    u.Name,
		Phone:   u.Phone,
		Address: u.Address,
	}
}
