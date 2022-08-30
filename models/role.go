package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model `json:"-"`
	ID         int    `json:"roleID" gorm:"primaryKey"`
	Name       string `json:"name"`
}
