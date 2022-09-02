package models

import "gorm.io/gorm"

type UserGift struct {
	gorm.Model `json:"-"`
	ID         int   `json:"userGiftID" gorm:"primaryKey"`
	UserID     int   `json:"userID"`
	User       *User `json:"user"`
	GiftID     int   `json:"giftID"`
	Gift       *Gift `json:"gift"`
	IsClaimed  int   `json:"isClaimed"`
}
