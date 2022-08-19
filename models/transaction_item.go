package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model     `json:"-"`
	ID             int          `json:"transactionItemID" gorm:"primaryKey"`
	SubscriptionID int          `json:"subscriptionID"`
	Subscription   Subscription `json:"-"`
	TransactionID  int          `json:"transactionID"`
	Transaction    Transaction  `json:"-"`
}
