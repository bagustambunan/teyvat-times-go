package models

import "gorm.io/gorm"

type TransactionStatus struct {
	gorm.Model `json:"-"`
	ID         int    `json:"transactionStatusID" gorm:"primaryKey"`
	Name       string `json:"name"`
}
