package models

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model  `json:"-"`
	ID          int    `json:"voucherID" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageID     int    `json:"imageID"`
	Image       Image  `json:"-"`
	Amount      int    `json:"amount"`
	Code        string `json:"code"`
}
