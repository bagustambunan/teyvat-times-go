package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model    `json:"-"`
	ID            int               `json:"transactionID" gorm:"primaryKey"`
	UserID        int               `json:"userID"`
	User          User              `json:"-"`
	StatusID      int               `json:"statusID"`
	Status        TransactionStatus `json:"-"`
	GrossTotal    int               `json:"grossTotal"`
	NetTotal      int               `json:"netTotal"`
	UserVoucherID int               `json:"userVoucherID"`
	UserVoucher   UserVoucher       `json:"-"`
}
