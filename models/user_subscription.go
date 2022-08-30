package models

import "gorm.io/gorm"

type UserSubscription struct {
	gorm.Model     `json:"-"`
	ID             int          `json:"userSubscriptionID" gorm:"primaryKey"`
	UserID         int          `json:"userID"`
	User           User         `json:"-"`
	SubscriptionID int          `json:"subscriptionID"`
	Subscription   Subscription `json:"-"`
	DateStart      string       `json:"dateStart"`
	DateEnded      string       `json:"dateEnded"`
}
