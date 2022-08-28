package models

import "gorm.io/gorm"

type UserVoucher struct {
	gorm.Model  `json:"-"`
	ID          int     `json:"userVoucherID" gorm:"primaryKey"`
	UserID      int     `json:"userID"`
	User        User    `json:"-"`
	VoucherID   int     `json:"-"`
	Voucher     Voucher `json:"voucher"`
	DateExpired string  `json:"dateExpired"`
	IsUsed      int     `json:"isUsed"`
}

func (uv *UserVoucher) GetVoucherAmount() int {
	return uv.Voucher.Amount
}
