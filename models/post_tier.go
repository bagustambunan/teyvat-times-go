package models

import "gorm.io/gorm"

type PostTier struct {
	gorm.Model    `json:"-"`
	ID            int    `json:"postTierID" gorm:"primaryKey"`
	Name          string `json:"name"`
	MoraRequired int    `json:"moraRequired"`
	Color         string `json:"color"`
}
