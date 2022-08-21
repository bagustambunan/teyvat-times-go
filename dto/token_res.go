package dto

type TokenRes struct {
	UserID int    `json:"userID"`
	Token  string `json:"token"`
}
