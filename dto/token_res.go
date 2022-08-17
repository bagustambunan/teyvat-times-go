package dto

type TokenRes struct {
	UserID  int    `json:"userID"`
	IDToken string `json:"idToken"`
}
