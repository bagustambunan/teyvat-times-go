package models

import "gorm.io/gorm"

type GiftClaimStatus struct {
	gorm.Model `json:"-"`
	ID         int    `json:"giftClaimStatusID" gorm:"primaryKey"`
	Name       string `json:"name"`
}
