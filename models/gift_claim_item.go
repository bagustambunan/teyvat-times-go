package models

import "gorm.io/gorm"

type GiftClaimItem struct {
	gorm.Model  `json:"-"`
	ID          int        `json:"giftClaimItemID" gorm:"primaryKey"`
	GiftID      int        `json:"giftID"`
	Gift        *Gift      `json:"gift"`
	GiftClaimID int        `json:"giftClaimID"`
	GiftClaim   *GiftClaim `json:"-"`
}
