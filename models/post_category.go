package models

import "gorm.io/gorm"

type PostCategory struct {
	gorm.Model `json:"-"`
	ID         int    `json:"postCategoryID" gorm:"primaryKey"`
	Name       string `json:"name"`
	Color      string `json:"color"`
}
