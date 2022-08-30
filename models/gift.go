package models

import "gorm.io/gorm"

type Gift struct {
	gorm.Model  `json:"-"`
	ID          int    `json:"giftID" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageID     int    `json:"-"`
	Image       *Image `json:"image"`
	Stock       int    `json:"stock"`
}
