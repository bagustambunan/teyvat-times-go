package models

import "gorm.io/gorm"

type UserReferral struct {
	gorm.Model     `json:"-"`
	ID             int `json:"userReferralID" gorm:"primaryKey"`
	UserID         int `json:"userID"`
	ReferrerUserID int `json:"referrerUserID"`
}
