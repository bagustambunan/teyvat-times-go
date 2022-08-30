package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model  `json:"-"`
	ID          int    `json:"subscriptionID" gorm:"primaryKey"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	MoraAmount int    `json:"moraAmount"`
}
