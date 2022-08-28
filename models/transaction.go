package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model     `json:"-"`
	ID             int               `json:"transactionID" gorm:"primaryKey"`
	UserID         int               `json:"-"`
	User           User              `json:"user"`
	SubscriptionID int               `json:"-"`
	Subscription   Subscription      `json:"subscription"`
	StatusID       int               `json:"-"`
	Status         TransactionStatus `json:"status"`
	GrossTotal     int               `json:"grossTotal"`
	NetTotal       int               `json:"netTotal"`
	UserVoucherID  int               `json:"userVoucherID"`
	UserVoucher    UserVoucher       `json:"-"`
	CreatedAt      string            `json:"createdAt"`
}
