package models

import "gorm.io/gorm"

type GiftClaim struct {
	gorm.Model     `json:"-"`
	ID             int              `json:"giftClaimID" gorm:"primaryKey"`
	UserID         int              `json:"userID"`
	User           *User            `json:"user"`
	AddressID      int              `json:"addressID"`
	Address        *Address         `json:"address"`
	StatusID       int              `json:"statusID"`
	Status         *GiftClaimStatus `json:"status"`
	GiftClaimItems []*GiftClaimItem `json:"giftClaimItems"`
}
