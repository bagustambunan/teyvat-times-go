package models

import "gorm.io/gorm"

type GiftClaim struct {
	gorm.Model     `json:"-"`
	ID             int              `json:"giftClaimID" gorm:"primaryKey"`
	UserID         int              `json:"userID"`
	User           *User            `json:"-"`
	AddressID      int              `json:"addressID"`
	Address        *Address         `json:"-"`
	StatusID       int              `json:"statusID"`
	Status         *GiftClaimStatus `json:"-"`
	GiftClaimItems []*GiftClaimItem `json:"giftClaimItems"`
}
