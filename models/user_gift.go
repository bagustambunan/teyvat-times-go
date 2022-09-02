package models

import "gorm.io/gorm"

type UserGift struct {
	gorm.Model `json:"-"`
	ID         int   `json:"userGiftID" gorm:"primaryKey"`
	UserID     int   `json:"userID"`
	User       *User `json:"-"`
	GiftID     int   `json:"giftID"`
	Gift       *Gift `json:"-"`
	IsClaimed  int   `json:"isClaimed"`
}
