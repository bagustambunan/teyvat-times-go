package models

import "gorm.io/gorm"

type ReferralCode struct {
	gorm.Model `json:"-"`
	ID         int    `json:"referralCodeID" gorm:"primaryKey"`
	Code       string `json:"code"`
}
