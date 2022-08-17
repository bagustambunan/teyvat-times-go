package dto

type SignUpReq struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	AddressStreet  string `json:"addressStreet"`
	AddressCity    string `json:"addressCity"`
	AddressState   string `json:"addressState"`
	AddressCountry string `json:"addressCountry"`
	PostalCode     string `json:"postalCode"`
	ReferrerCode   string `json:"referrerCode"`
}
