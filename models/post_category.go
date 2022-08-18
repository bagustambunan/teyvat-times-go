package models

import "gorm.io/gorm"

type PostCategory struct {
	gorm.Model `json:"-"`
	ID         int    `json:"roleID" gorm:"primaryKey"`
	Name       string `json:"name"`
}
