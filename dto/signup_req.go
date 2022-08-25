package dto

type SignUpReq struct {
	Name         string `json:"name" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	PostalCode   string `json:"postalCode" binding:"required"`
	ReferrerCode string `json:"referrerCode"`
}
