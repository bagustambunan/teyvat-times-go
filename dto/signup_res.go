package dto

import "git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"

type SignUpRes struct {
	ID int `json:"userID"`
	//Username string          `json:"username"`
	Email   string          `json:"email"`
	Name    string          `json:"name"`
	Phone   string          `json:"phone"`
	Address *models.Address `json:"address"`
}

func (_ *SignUpRes) FromUser(u *models.User) *SignUpRes {
	return &SignUpRes{
		ID: u.ID,
		//Username: u.Username,
		Email:   u.Email,
		Name:    u.Name,
		Phone:   u.Phone,
		Address: u.Address,
	}
}
