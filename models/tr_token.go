package models

import "gorm.io/gorm"

type TrToken struct {
	gorm.Model `json:"-"`
	ID         int    `json:"trTokenID" gorm:"primaryKey"`
	UserID     int    `json:"userID"`
	User       User   `json:"-"`
	Token      string `json:"token"`
	IsExpired  int    `json:"isExpired"`
}
