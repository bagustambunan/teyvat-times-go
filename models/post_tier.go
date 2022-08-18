package models

import "gorm.io/gorm"

type PostTier struct {
	gorm.Model    `json:"-"`
	ID            int    `json:"roleID" gorm:"primaryKey"`
	Name          string `json:"name"`
	CoinsRequired int    `json:"coinsRequired"`
}
